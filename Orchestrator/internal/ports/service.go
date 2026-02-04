package ports

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
)

// OrchestratorService defines the interface for the Orchestrator business logic
type OrchestratorService interface {
	RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	UpgradeSubscription(ctx context.Context, req *dto.UpgradeSubscriptionRequest) (*dto.UpgradeSubscriptionResponse, error)
}
