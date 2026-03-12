package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
)

// ToUserPreferenceModel converts a domain entity to a DB model.
func ToUserPreferenceModel(e *entity.UserPreference) *models.UserPreference {
	if e == nil {
		return nil
	}

	id := primitive.NewObjectID()
	if e.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(e.ID); err == nil {
			id = oid
		}
	}

	return &models.UserPreference{
		BaseModel: &mongodb.BaseModel{
			ID:        id,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		UserID:          e.UserID,
		EmailEnabled:    e.EmailEnabled,
		InAppEnabled:    e.InAppEnabled,
		WebhookEnabled:  e.WebhookEnabled,
		QuietHoursStart: e.QuietHoursStart,
		QuietHoursEnd:   e.QuietHoursEnd,
		Timezone:        e.Timezone,
	}
}

// ToUserPreferenceEntity converts a DB model back to a domain entity.
func ToUserPreferenceEntity(m *models.UserPreference) *entity.UserPreference {
	if m == nil {
		return nil
	}

	return &entity.UserPreference{
		ID:              m.ID.Hex(),
		UserID:          m.UserID,
		EmailEnabled:    m.EmailEnabled,
		InAppEnabled:    m.InAppEnabled,
		WebhookEnabled:  m.WebhookEnabled,
		QuietHoursStart: m.QuietHoursStart,
		QuietHoursEnd:   m.QuietHoursEnd,
		Timezone:        m.Timezone,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}
