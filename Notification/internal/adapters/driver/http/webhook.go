package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/ports"
)

// WebhookConfigHandler defines the HTTP handler interface for webhook configurations.
type WebhookConfigHandler interface {
	Create(ctx context.Context, req *dto.CreateWebhookRequest) (*dto.WebhookResponse, error)
	Get(ctx context.Context, req *dto.GetWebhookRequest) (*dto.WebhookResponse, error)
	List(ctx context.Context, _ *struct{}) ([]*dto.WebhookResponse, error)
	Update(ctx context.Context, req *dto.UpdateWebhookRequest) (*dto.WebhookResponse, error)
	Delete(ctx context.Context, req *dto.DeleteWebhookRequest) (*struct{}, error)
}

type webhookConfigHandler struct {
	handler.BaseHandler
	service ports.WebhookConfigService
}

// NewWebhookConfigHandler creates a new WebhookConfigHandler.
func NewWebhookConfigHandler(service ports.WebhookConfigService) WebhookConfigHandler {
	return &webhookConfigHandler{service: service}
}

// Create adds a new webhook configuration.
func (h *webhookConfigHandler) Create(ctx context.Context, req *dto.CreateWebhookRequest) (*dto.WebhookResponse, error) {
	return h.service.Create(ctx, req)
}

// Get retrieves a webhook configuration by ID.
func (h *webhookConfigHandler) Get(ctx context.Context, req *dto.GetWebhookRequest) (*dto.WebhookResponse, error) {
	return h.service.Get(ctx, req.ID)
}

// List retrieves all webhook configurations for the tenant.
func (h *webhookConfigHandler) List(ctx context.Context, _ *struct{}) ([]*dto.WebhookResponse, error) {
	return h.service.FindByTenantID(ctx)
}

// Update modifies an existing webhook configuration.
func (h *webhookConfigHandler) Update(ctx context.Context, req *dto.UpdateWebhookRequest) (*dto.WebhookResponse, error) {
	return h.service.Update(ctx, req)
}

// Delete removes a webhook configuration.
func (h *webhookConfigHandler) Delete(ctx context.Context, req *dto.DeleteWebhookRequest) (*struct{}, error) {
	return &struct{}{}, h.service.Delete(ctx, req.ID)
}
