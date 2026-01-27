package ports

import (
	"context"

	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
)

// PlanRepository defines the plan data access interface.
type PlanRepository interface {
	Get(ctx context.Context, id int) (*entity.Plan, error)
	Create(ctx context.Context, e *entity.Plan) error
	Update(ctx context.Context, e *entity.Plan) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*entity.Plan, error)
	FindActive(ctx context.Context) ([]*entity.Plan, error)
}

// PlanService defines the plan business logic interface.
type PlanService interface {
	Get(ctx context.Context, id int) (*dto.PlanResponse, error)
	Create(ctx context.Context, req *dto.CreatePlanRequest) (*dto.PlanResponse, error)
	Update(ctx context.Context, id int, req *dto.UpdatePlanRequest) (*dto.PlanResponse, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*dto.PlanResponse, error)
	FindActive(ctx context.Context) ([]*dto.PlanResponse, error)
}
