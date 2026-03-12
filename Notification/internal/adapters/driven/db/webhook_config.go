package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/mapper"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
)

const (
	WebhookConfigCollection = "webhook_configs"
)

type webhookConfigRepo struct {
	repo *mongodb.BaseRepository[models.WebhookConfig]
}

// NewWebhookConfigRepo creates a MongoDB-backed WebhookConfigRepository.
func NewWebhookConfigRepo(db *mongo.Database) ports.WebhookConfigRepository {
	collection := db.Collection(WebhookConfigCollection)
	return &webhookConfigRepo{
		repo: mongodb.NewBaseRepository[models.WebhookConfig](collection),
	}
}

// Create inserts a new webhook config.
func (r *webhookConfigRepo) Create(ctx context.Context, config *entity.WebhookConfig) error {
	model := mapper.ToWebhookConfigModel(config)

	err := r.repo.Create(ctx, model)
	if err != nil {
		return err
	}

	config.ID = model.ID.Hex()
	return nil
}

// Get retrieves a webhook config by hex ID string.
func (r *webhookConfigRepo) Get(ctx context.Context, id string) (*entity.WebhookConfig, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	model, err := r.repo.Get(ctx, oid)
	if err != nil {
		return nil, err
	}

	return mapper.ToWebhookConfigEntity(model), nil
}

// GetByTenantID returns all webhook configs for the given tenant.
func (r *webhookConfigRepo) GetByTenantID(ctx context.Context, tenantID string) ([]*entity.WebhookConfig, error) {
	// Fallback to native Find since BaseRepo Find uses complex QueryOptions DTO.
	// For simple queries, using the underlying collection is sometimes more straightforward.
	cursor, err := r.repo.GetCollection().Find(ctx, bson.M{"tenant_id": tenantID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var modelsList []*models.WebhookConfig
	if err := cursor.All(ctx, &modelsList); err != nil {
		return nil, err
	}

	configs := make([]*entity.WebhookConfig, 0, len(modelsList))
	for _, m := range modelsList {
		configs = append(configs, mapper.ToWebhookConfigEntity(m))
	}
	return configs, nil
}

// Update replaces mutable fields on an existing webhook config.
func (r *webhookConfigRepo) Update(ctx context.Context, config *entity.WebhookConfig) error {
	model := mapper.ToWebhookConfigModel(config)
	return r.repo.Update(ctx, model)
}

// Delete removes a webhook config by hex ID string.
func (r *webhookConfigRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return r.repo.Delete(ctx, oid)
}
