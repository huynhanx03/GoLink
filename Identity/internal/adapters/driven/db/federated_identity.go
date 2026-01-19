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

type FederatedIdentityRepository struct {
	repo   *ent.BaseRepository[generate.FederatedIdentity, *generate.FederatedIdentity, int]
	client *generate.FederatedIdentityClient
}

// NewFederatedIdentityRepository creates a new FederatedIdentityRepository instance.
func NewFederatedIdentityRepository(client interface{}) ports.FederatedIdentityRepository {
	entClient := client.(*generate.Client)
	return &FederatedIdentityRepository{
		repo:   ent.NewBaseRepository[generate.FederatedIdentity, *generate.FederatedIdentity, int](client),
		client: entClient.FederatedIdentity,
	}
}

// Find retrieves federated identities with pagination.
func (r *FederatedIdentityRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.FederatedIdentity], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.FederatedIdentity, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToFederatedIdentityEntity(record)
	}

	return &d.Paginated[*entity.FederatedIdentity]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a federated identity by ID.
func (r *FederatedIdentityRepository) Get(ctx context.Context, id int) (*entity.FederatedIdentity, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToFederatedIdentityEntity(record), nil
}

// Create creates a new federated identity.
func (r *FederatedIdentityRepository) Create(ctx context.Context, e *entity.FederatedIdentity) error {
	model := mapper.ToFederatedIdentityModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToFederatedIdentityEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Delete removes a federated identity by ID.
func (r *FederatedIdentityRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a federated identity exists by ID.
func (r *FederatedIdentityRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}
