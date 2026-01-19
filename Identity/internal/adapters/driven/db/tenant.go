package db

import (
	"context"

	"go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

type TenantRepository struct {
	repo   *ent.BaseRepository[generate.Tenant, *generate.Tenant, int]
	client *generate.TenantClient
}

// NewTenantRepository creates a new TenantRepository instance.
func NewTenantRepository(client interface{}) ports.TenantRepository {
	entClient := client.(*generate.Client)
	return &TenantRepository{
		repo:   ent.NewBaseRepository[generate.Tenant, *generate.Tenant, int](client),
		client: entClient.Tenant,
	}
}

func (r *TenantRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Tenant], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Tenant, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToTenantEntity(record)
	}

	return &d.Paginated[*entity.Tenant]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a tenant by ID.
func (r *TenantRepository) Get(ctx context.Context, id int) (*entity.Tenant, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToTenantEntity(record), nil
}

// Create creates a new tenant.
func (r *TenantRepository) Create(ctx context.Context, e *entity.Tenant) error {
	model := mapper.ToTenantModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToTenantEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing tenant.
func (r *TenantRepository) Update(ctx context.Context, e *entity.Tenant) error {
	model := mapper.ToTenantModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a tenant by ID.
func (r *TenantRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a tenant exists by ID.
func (r *TenantRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}
