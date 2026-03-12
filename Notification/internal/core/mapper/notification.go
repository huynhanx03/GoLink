package mapper

import (
	"time"

	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
)

// ToNotificationResponse converts a domain entity to a response DTO.
func ToNotificationResponse(e *entity.Notification) *dto.NotificationResponse {
	if e == nil {
		return nil
	}

	var sentAt *string
	if e.SentAt != nil {
		s := e.SentAt.Format(time.RFC3339)
		sentAt = &s
	}

	return &dto.NotificationResponse{
		ID:           e.ID,
		Type:         e.Type,
		Channel:      e.Channel,
		Priority:     e.Priority,
		Status:       e.Status,
		Subject:      e.Subject,
		Body:         e.Body,
		TemplateData: e.TemplateData,
		CollapseKey:  e.CollapseKey,
		IsRead:       e.IsRead,
		ErrorMessage: e.ErrorMessage,
		CreatedAt:    e.CreatedAt.Format(time.RFC3339),
		SentAt:       sentAt,
	}
}

// ToNotificationResponseList converts a list of domain entities to a list of response DTOs.
func ToNotificationResponseList(list []*entity.Notification) []*dto.NotificationResponse {
	if list == nil {
		return nil
	}

	res := make([]*dto.NotificationResponse, 0, len(list))
	for _, e := range list {
		res = append(res, ToNotificationResponse(e))
	}
	return res
}
