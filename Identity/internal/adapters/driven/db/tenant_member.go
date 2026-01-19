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

type TenantMemberRepository struct {
	repo   *ent.BaseRepository[generate.TenantMember, *generate.TenantMember, int]
	client *generate.TenantMemberClient
}

// NewTenantMemberRepository creates a new TenantMemberRepository instance.
func NewTenantMemberRepository(client interface{}) ports.TenantMemberRepository {
	entClient := client.(*generate.Client)
	return &TenantMemberRepository{
		repo:   ent.NewBaseRepository[generate.TenantMember, *generate.TenantMember, int](client),
		client: entClient.TenantMember,
	}
}

// Find retrieves tenant members with pagination.
func (r *TenantMemberRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.TenantMember], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.TenantMember, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToTenantMemberEntity(record)
	}

	return &d.Paginated[*entity.TenantMember]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a tenant member by ID.
func (r *TenantMemberRepository) Get(ctx context.Context, id int) (*entity.TenantMember, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToTenantMemberEntity(record), nil
}

// Create creates a new tenant member.
func (r *TenantMemberRepository) Create(ctx context.Context, e *entity.TenantMember) error {
	model := mapper.ToTenantMemberModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToTenantMemberEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing tenant member.
func (r *TenantMemberRepository) Update(ctx context.Context, e *entity.TenantMember) error {
	model := mapper.ToTenantMemberModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a tenant member by ID.
func (r *TenantMemberRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a tenant member exists by ID.
func (r *TenantMemberRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}
