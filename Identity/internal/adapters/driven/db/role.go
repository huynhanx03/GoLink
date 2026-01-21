package db

import (
	"context"

	"go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driven/db/ent/generate/role"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

type RoleRepository struct {
	repo   *ent.BaseRepository[generate.Role, *generate.Role, int]
	client *generate.RoleClient
}

// NewRoleRepository creates a new RoleRepository instance.
func NewRoleRepository(client interface{}) ports.RoleRepository {
	entClient := client.(*generate.Client)
	return &RoleRepository{
		repo:   ent.NewBaseRepository[generate.Role, *generate.Role, int](client),
		client: entClient.Role,
	}
}

// FindAll retrieves all roles.
func (r *RoleRepository) FindAll(ctx context.Context) ([]*entity.Role, error) {
	models, err := r.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Role, len(models))
	for i, m := range models {
		entities[i] = mapper.ToRoleEntity(m)
	}
	return entities, nil
}

// Find retrieves roles with pagination.
func (r *RoleRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Role], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Role, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToRoleEntity(record)
	}

	return &d.Paginated[*entity.Role]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a role by ID.
func (r *RoleRepository) Get(ctx context.Context, id int) (*entity.Role, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToRoleEntity(record), nil
}

// GetByName retrieves a role by name.
func (r *RoleRepository) GetByName(ctx context.Context, name string) (*entity.Role, error) {
	record, err := r.client.Query().
		Where(role.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToRoleEntity(record), nil
}

// Create creates a new role.
func (r *RoleRepository) Create(ctx context.Context, e *entity.Role) error {
	model := mapper.ToRoleModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToRoleEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing role.
func (r *RoleRepository) Update(ctx context.Context, e *entity.Role) error {
	model := mapper.ToRoleModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// UpdateBulk updates multiple roles.
func (r *RoleRepository) UpdateBulk(ctx context.Context, entities []*entity.Role) error {
	models := make([]*generate.Role, len(entities))
	for i, e := range entities {
		models[i] = mapper.ToRoleModel(e)
	}

	if err := r.repo.UpdateBulk(ctx, models); err != nil {
		return err
	}

	for i, m := range models {
		entities[i].UpdatedAt = m.UpdatedAt
	}

	return nil
}

// Delete removes a role by ID.
func (r *RoleRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a role exists by ID.
func (r *RoleRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}

// FindDescendants retrieves all descendant roles of a given role (inclusive).
func (r *RoleRepository) FindDescendants(ctx context.Context, lft, rgt int) ([]*entity.Role, error) {
	models, err := r.client.Query().
		Where(role.LftGTE(lft), role.RgtLTE(rgt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Role, len(models))
	for i, m := range models {
		entities[i] = mapper.ToRoleEntity(m)
	}
	return entities, nil
}
