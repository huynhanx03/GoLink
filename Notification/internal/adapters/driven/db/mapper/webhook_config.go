package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
)

// ToWebhookConfigModel converts a domain entity to a DB model.
func ToWebhookConfigModel(e *entity.WebhookConfig) *models.WebhookConfig {
	if e == nil {
		return nil
	}

	id := primitive.NewObjectID()
	if e.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(e.ID); err == nil {
			id = oid
		}
	}

	return &models.WebhookConfig{
		BaseModel: &mongodb.BaseModel{
			ID:        id,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		TenantID:   e.TenantID,
		URL:        e.URL,
		Secret:     e.Secret,
		EventTypes: e.EventTypes,
		IsActive:   e.IsActive,
	}
}

// ToWebhookConfigEntity converts a DB model back to a domain entity.
func ToWebhookConfigEntity(m *models.WebhookConfig) *entity.WebhookConfig {
	if m == nil {
		return nil
	}

	return &entity.WebhookConfig{
		ID:         m.ID.Hex(),
		TenantID:   m.TenantID,
		URL:        m.URL,
		Secret:     m.Secret,
		EventTypes: m.EventTypes,
		IsActive:   m.IsActive,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}
