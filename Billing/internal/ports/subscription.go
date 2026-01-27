package ports

import (
	"context"

	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
)

// SubscriptionRepository defines the subscription data access interface.
type SubscriptionRepository interface {
	Get(ctx context.Context, id int) (*entity.Subscription, error)
	Create(ctx context.Context, e *entity.Subscription) error
	Update(ctx context.Context, e *entity.Subscription) error
	Delete(ctx context.Context, id int) error
}

// SubscriptionService defines the subscription business logic interface.
type SubscriptionService interface {
	Get(ctx context.Context, id int) (*dto.SubscriptionResponse, error)
	Create(ctx context.Context, req *dto.CreateSubscriptionRequest) (*dto.SubscriptionResponse, error)
	Update(ctx context.Context, id int, req *dto.UpdateSubscriptionRequest) (*dto.SubscriptionResponse, error)
	Delete(ctx context.Context, id int) error
}
