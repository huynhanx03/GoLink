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

type ResourceRepository struct {
	repo   *ent.BaseRepository[generate.Resource, *generate.Resource, int]
	client *generate.ResourceClient
}

// NewResourceRepository creates a new ResourceRepository instance.
func NewResourceRepository(client interface{}) ports.ResourceRepository {
	entClient := client.(*generate.Client)
	return &ResourceRepository{
		repo:   ent.NewBaseRepository[generate.Resource, *generate.Resource, int](client),
		client: entClient.Resource,
	}
}

// Find retrieves resources with pagination.
func (r *ResourceRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Resource], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Resource, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToResourceEntity(record)
	}

	return &d.Paginated[*entity.Resource]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a resource by ID.
func (r *ResourceRepository) Get(ctx context.Context, id int) (*entity.Resource, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToResourceEntity(record), nil
}

// Create creates a new resource.
func (r *ResourceRepository) Create(ctx context.Context, e *entity.Resource) error {
	model := mapper.ToResourceModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToResourceEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing resource.
func (r *ResourceRepository) Update(ctx context.Context, e *entity.Resource) error {
	model := mapper.ToResourceModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a resource by ID.
func (r *ResourceRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a resource exists by ID.
func (r *ResourceRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}

// FindByIDs retrieves resources by a list of IDs.
func (r *ResourceRepository) FindByIDs(ctx context.Context, ids []int) ([]*entity.Resource, error) {
	models, err := r.repo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Resource, len(models))
	for i, m := range models {
		entities[i] = mapper.ToResourceEntity(m)
	}
	return entities, nil
}
