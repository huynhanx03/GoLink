package db

import (
	"context"

	"go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driven/db/ent/generate/permission"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

type PermissionRepository struct {
	repo   *ent.BaseRepository[generate.Permission, *generate.Permission, int]
	client *generate.PermissionClient
}

// NewPermissionRepository creates a new PermissionRepository instance.
func NewPermissionRepository(client interface{}) ports.PermissionRepository {
	entClient := client.(*generate.Client)
	return &PermissionRepository{
		repo:   ent.NewBaseRepository[generate.Permission, *generate.Permission, int](client),
		client: entClient.Permission,
	}
}

// Find retrieves permissions with pagination.
func (r *PermissionRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Permission], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Permission, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToPermissionEntity(record)
	}

	return &d.Paginated[*entity.Permission]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a permission by ID.
func (r *PermissionRepository) Get(ctx context.Context, id int) (*entity.Permission, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToPermissionEntity(record), nil
}

// Create creates a new permission.
func (r *PermissionRepository) Create(ctx context.Context, e *entity.Permission) error {
	model := mapper.ToPermissionModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToPermissionEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing permission.
func (r *PermissionRepository) Update(ctx context.Context, e *entity.Permission) error {
	model := mapper.ToPermissionModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a permission by ID.
func (r *PermissionRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a permission exists by ID.
func (r *PermissionRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}

// FindByRoleIDs retrieves permissions for a list of role IDs.
func (r *PermissionRepository) FindByRoleIDs(ctx context.Context, roleIDs []int) ([]*entity.Permission, error) {
	models, err := r.client.Query().
		Where(permission.RoleIDIn(roleIDs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Permission, len(models))
	for i, m := range models {
		entities[i] = mapper.ToPermissionEntity(m)
	}
	return entities, nil
}
