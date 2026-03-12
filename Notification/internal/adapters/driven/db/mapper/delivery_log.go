package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-link/common/pkg/database/mongodb"
	"go-link/notification/internal/adapters/driven/db/models"
	"go-link/notification/internal/core/entity"
)

// ToDeliveryLogModel converts a domain entity to a DB model.
func ToDeliveryLogModel(e *entity.DeliveryLog) *models.DeliveryLog {
	if e == nil {
		return nil
	}

	id := primitive.NewObjectID()
	if e.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(e.ID); err == nil {
			id = oid
		}
	}

	return &models.DeliveryLog{
		BaseModel: &mongodb.BaseModel{
			ID:        id,
			CreatedAt: e.CreatedAt,
		},
		NotificationID: e.NotificationID,
		Channel:        e.Channel,
		Status:         e.Status,
		Attempt:        e.Attempt,
		ErrorMessage:   e.ErrorMessage,
		ResponseCode:   e.ResponseCode,
		Duration:       e.Duration,
	}
}

// ToDeliveryLogEntity converts a DB model back to a domain entity.
func ToDeliveryLogEntity(m *models.DeliveryLog) *entity.DeliveryLog {
	if m == nil {
		return nil
	}

	return &entity.DeliveryLog{
		ID:             m.ID.Hex(),
		NotificationID: m.NotificationID,
		Channel:        m.Channel,
		Status:         m.Status,
		Attempt:        m.Attempt,
		ErrorMessage:   m.ErrorMessage,
		ResponseCode:   m.ResponseCode,
		Duration:       m.Duration,
		CreatedAt:      m.CreatedAt,
	}
}
