package db

import (
	"context"

	"go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driven/db/ent/generate/user"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

type UserRepository struct {
	repo   *ent.BaseRepository[generate.User, *generate.User, int]
	client *generate.UserClient
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(client interface{}) ports.UserRepository {
	entClient := client.(*generate.Client)
	return &UserRepository{
		repo:   ent.NewBaseRepository[generate.User, *generate.User, int](client),
		client: entClient.User,
	}
}

// Find retrieves users with pagination.
func (r *UserRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.User], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.User, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToUserEntity(record)
	}

	return &d.Paginated[*entity.User]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves a user by ID.
func (r *UserRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(record), nil
}

// GetByUsername retrieves a user by username.
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	record, err := r.client.Query().
		Where(user.Username(username)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(record), nil
}

// Create creates a new user.
func (r *UserRepository) Create(ctx context.Context, e *entity.User) error {
	model := mapper.ToUserModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToUserEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing user.
func (r *UserRepository) Update(ctx context.Context, e *entity.User) error {
	model := mapper.ToUserModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a user by ID.
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if a user exists by ID.
func (r *UserRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}

// ExistsByUsername checks if a user exists by username.
func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	return r.client.Query().Where(user.Username(username)).Exist(ctx)
}
