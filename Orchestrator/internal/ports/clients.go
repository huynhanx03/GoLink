package ports

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
)

// IdentityClient defines the interface for interacting with Identity Service
type IdentityClient interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error)
	DeleteUser(ctx context.Context, userID int64) error
}

// BillingClient defines the interface for interacting with Billing Service
type BillingClient interface {
	CreateSubscription(ctx context.Context, userID int64, planID int) (int64, error)
	CancelSubscription(ctx context.Context, subscriptionID int64) error
}
