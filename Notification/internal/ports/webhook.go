package ports

import (
	"context"

	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/entity"
)

// WebhookConfigRepository defines persistence operations for tenant webhook configurations.
type WebhookConfigRepository interface {
	Create(ctx context.Context, config *entity.WebhookConfig) error
	Get(ctx context.Context, id string) (*entity.WebhookConfig, error)
	GetByTenantID(ctx context.Context, tenantID string) ([]*entity.WebhookConfig, error)
	Update(ctx context.Context, config *entity.WebhookConfig) error
	Delete(ctx context.Context, id string) error
}

// WebhookConfigService defines the business logic contract for webhook configurations.
type WebhookConfigService interface {
	Create(ctx context.Context, req *dto.CreateWebhookRequest) (*dto.WebhookResponse, error)
	Get(ctx context.Context, id string) (*dto.WebhookResponse, error)
	FindByTenantID(ctx context.Context) ([]*dto.WebhookResponse, error)
	Update(ctx context.Context, req *dto.UpdateWebhookRequest) (*dto.WebhookResponse, error)
	Delete(ctx context.Context, id string) error
}
