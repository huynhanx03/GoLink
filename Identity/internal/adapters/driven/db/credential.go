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

type CredentialRepository struct {
	repo   *ent.BaseRepository[generate.Credential, *generate.Credential, int]
	client *generate.CredentialClient
}

// NewCredentialRepository creates a new CredentialRepository instance.
func NewCredentialRepository(client interface{}) ports.CredentialRepository {
	entClient := client.(*generate.Client)
	return &CredentialRepository{
		repo:   ent.NewBaseRepository[generate.Credential, *generate.Credential, int](client),
		client: entClient.Credential,
	}
}

// Find retrieves credentials with pagination.
func (r *CredentialRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Credential], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Credential, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToCredentialEntity(record)
	}

	return &d.Paginated[*entity.Credential]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a credential by ID.
func (r *CredentialRepository) Get(ctx context.Context, id int) (*entity.Credential, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToCredentialEntity(record), nil
}

// Create creates a new credential.
func (r *CredentialRepository) Create(ctx context.Context, e *entity.Credential) error {
	model := mapper.ToCredentialModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToCredentialEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing credential.
func (r *CredentialRepository) Update(ctx context.Context, e *entity.Credential) error {
	model := mapper.ToCredentialModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a credential by ID.
func (r *CredentialRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a credential exists by ID.
func (r *CredentialRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}
