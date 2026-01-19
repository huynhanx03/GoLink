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

type DomainRepository struct {
	repo   *ent.BaseRepository[generate.Domain, *generate.Domain, int]
	client *generate.DomainClient
}

// NewDomainRepository creates a new DomainRepository instance.
func NewDomainRepository(client interface{}) ports.DomainRepository {
	entClient := client.(*generate.Client)
	return &DomainRepository{
		repo:   ent.NewBaseRepository[generate.Domain, *generate.Domain, int](client),
		client: entClient.Domain,
	}
}

// Find retrieves domains with pagination.
func (r *DomainRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Domain], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Domain, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToDomainEntity(record)
	}

	return &d.Paginated[*entity.Domain]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a domain by ID.
func (r *DomainRepository) Get(ctx context.Context, id int) (*entity.Domain, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToDomainEntity(record), nil
}

// Create creates a new domain.
func (r *DomainRepository) Create(ctx context.Context, e *entity.Domain) error {
	model := mapper.ToDomainModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToDomainEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing domain.
func (r *DomainRepository) Update(ctx context.Context, e *entity.Domain) error {
	model := mapper.ToDomainModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a domain by ID.
func (r *DomainRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a domain exists by ID.
func (r *DomainRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}
