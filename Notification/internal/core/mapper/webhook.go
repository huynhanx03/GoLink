package mapper

import (
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
)

// ToWebhookEntityFromCreate converts a create request DTO to a domain entity.
func ToWebhookEntityFromCreate(tenantID string, req *dto.CreateWebhookRequest) *entity.WebhookConfig {
	if req == nil {
		return nil
	}

	return &entity.WebhookConfig{
		TenantID:   tenantID,
		URL:        req.URL,
		EventTypes: req.EventTypes,
		IsActive:   req.IsActive,
	}
}

// ToWebhookEntityFromUpdate converts an update request DTO to a domain entity.
func ToWebhookEntityFromUpdate(tenantID, id string, req *dto.UpdateWebhookRequest) *entity.WebhookConfig {
	if req == nil {
		return nil
	}

	return &entity.WebhookConfig{
		ID:         id,
		TenantID:   tenantID,
		URL:        req.URL,
		EventTypes: req.EventTypes,
		IsActive:   req.IsActive,
	}
}

// ToWebhookResponse converts a domain entity to a response DTO.
func ToWebhookResponse(e *entity.WebhookConfig) *dto.WebhookResponse {
	if e == nil {
		return nil
	}

	return &dto.WebhookResponse{
		ID:         e.ID,
		TenantID:   e.TenantID,
		URL:        e.URL,
		EventTypes: e.EventTypes,
		IsActive:   e.IsActive,
	}
}

// ToWebhookResponseList converts a list of domain entities to a list of response DTOs.
func ToWebhookResponseList(list []*entity.WebhookConfig) []*dto.WebhookResponse {
	if list == nil {
		return nil
	}

	res := make([]*dto.WebhookResponse, 0, len(list))
	for _, e := range list {
		res = append(res, ToWebhookResponse(e))
	}
	return res
}
