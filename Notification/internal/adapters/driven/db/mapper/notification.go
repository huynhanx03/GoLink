package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
)

// ToNotificationModel converts a domain entity to a DB model.
func ToNotificationModel(e *entity.Notification) *models.Notification {
	if e == nil {
		return nil
	}

	id := primitive.NewObjectID()
	if e.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(e.ID); err == nil {
			id = oid
		}
	}

	return &models.Notification{
		BaseModel: &mongodb.BaseModel{
			ID:        id,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		IdempotencyKey: e.IdempotencyKey,
		Type:           e.Type,
		Channel:        e.Channel,
		Priority:       e.Priority,
		Status:         e.Status,
		Recipient: models.Recipient{
			UserID: e.Recipient.UserID,
			Email:  e.Recipient.Email,
			Name:   e.Recipient.Name,
		},
		Subject:      e.Subject,
		Body:         e.Body,
		TemplateData: e.TemplateData,
		CollapseKey:  e.CollapseKey,
		IsRead:       e.IsRead,
		ErrorMessage: e.ErrorMessage,
		RetryCount:   e.RetryCount,
		SentAt:       e.SentAt,
		ExpiresAt:    e.ExpiresAt,
	}
}

// ToNotificationEntity converts a DB model back to a domain entity.
func ToNotificationEntity(m *models.Notification) *entity.Notification {
	if m == nil {
		return nil
	}

	return &entity.Notification{
		ID:             m.ID.Hex(),
		IdempotencyKey: m.IdempotencyKey,
		Type:           m.Type,
		Channel:        m.Channel,
		Priority:       m.Priority,
		Status:         m.Status,
		Recipient: entity.Recipient{
			UserID: m.Recipient.UserID,
			Email:  m.Recipient.Email,
			Name:   m.Recipient.Name,
		},
		Subject:      m.Subject,
		Body:         m.Body,
		TemplateData: m.TemplateData,
		CollapseKey:  m.CollapseKey,
		IsRead:       m.IsRead,
		ErrorMessage: m.ErrorMessage,
		RetryCount:   m.RetryCount,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		SentAt:       m.SentAt,
		ExpiresAt:    m.ExpiresAt,
	}
}
