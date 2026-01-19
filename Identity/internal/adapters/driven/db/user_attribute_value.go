package db

import (
	"context"

	"go-link/common/pkg/database/ent"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driven/db/ent/generate/userattributevalue"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

type UserAttributeValueRepository struct {
	repo   *ent.BaseRepository[generate.UserAttributeValue, *generate.UserAttributeValue, int]
	client *generate.UserAttributeValueClient
}

// NewUserAttributeValueRepository creates a new UserAttributeValueRepository instance.
func NewUserAttributeValueRepository(client interface{}) ports.UserAttributeValueRepository {
	entClient := client.(*generate.Client)
	return &UserAttributeValueRepository{
		repo:   ent.NewBaseRepository[generate.UserAttributeValue, *generate.UserAttributeValue, int](client),
		client: entClient.UserAttributeValue,
	}
}

// Get retrieves a user attribute value by ID.
func (r *UserAttributeValueRepository) Get(ctx context.Context, id int) (*entity.UserAttributeValue, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserAttributeValueEntity(record), nil
}

// GetByUserID retrieves all attribute values for a user.
func (r *UserAttributeValueRepository) GetByUserID(ctx context.Context, userID int) ([]*entity.UserAttributeValue, error) {
	records, err := r.client.Query().
		Where(userattributevalue.UserID(userID)).
		WithDefinition().
		All(ctx)
	if err != nil {
		return nil, err
	}
	entities := make([]*entity.UserAttributeValue, len(records))
	for i, record := range records {
		entities[i] = mapper.ToUserAttributeValueEntity(record)
	}
	return entities, nil
}

// Create creates a new user attribute value.
func (r *UserAttributeValueRepository) Create(ctx context.Context, e *entity.UserAttributeValue) error {
	model := mapper.ToUserAttributeValueModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToUserAttributeValueEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing user attribute value.
func (r *UserAttributeValueRepository) Update(ctx context.Context, e *entity.UserAttributeValue) error {
	model := mapper.ToUserAttributeValueModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes a user attribute value by ID.
func (r *UserAttributeValueRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// DeleteByUserID removes all attribute values for a user.
func (r *UserAttributeValueRepository) DeleteByUserID(ctx context.Context, userID int) error {
	_, err := r.client.Delete().Where(userattributevalue.UserID(userID)).Exec(ctx)
	return err
}
