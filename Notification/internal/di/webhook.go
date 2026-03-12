package di

import (
	httpAdapter "go-link/notification/internal/adapters/driver/http"
	"go-link/notification/internal/core/service"
	"go-link/notification/internal/ports"
)

// WebhookContainer holds all webhook-domain dependencies.
type WebhookContainer struct {
	Service ports.WebhookConfigService
	Handler httpAdapter.WebhookConfigHandler
}

// InitWebhookDependencies initializes dependencies for the Webhook domain.
func InitWebhookDependencies(
	repo ports.WebhookConfigRepository,
) WebhookContainer {
	svc := service.NewWebhookConfigService(repo)
	handler := httpAdapter.NewWebhookConfigHandler(svc)

	return WebhookContainer{
		Service: svc,
		Handler: handler,
	}
}
