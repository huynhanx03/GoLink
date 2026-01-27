package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"

	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/ports"
)

// PlanHandler defines the plan HTTP handler interface.
type PlanHandler interface {
	Get(ctx context.Context, req *dto.GetPlanRequest) (*dto.PlanResponse, error)
	Create(ctx context.Context, req *dto.CreatePlanRequest) (*dto.PlanResponse, error)
	Update(ctx context.Context, req *dto.UpdatePlanRequest) (*dto.PlanResponse, error)
	Delete(ctx context.Context, req *dto.DeletePlanRequest) (*dto.PlanResponse, error)
	FindAll(ctx context.Context, req *dto.FindPlanRequest) ([]*dto.PlanResponse, error)
	FindActive(ctx context.Context, req *dto.FindActivePlanRequest) ([]*dto.PlanResponse, error)
}

type planHandler struct {
	handler.BaseHandler
	planService ports.PlanService
}

// NewPlanHandler creates a new PlanHandler instance.
func NewPlanHandler(planService ports.PlanService) PlanHandler {
	return &planHandler{
		planService: planService,
	}
}

// Get retrieves a plan by ID.
func (h *planHandler) Get(ctx context.Context, req *dto.GetPlanRequest) (*dto.PlanResponse, error) {
	return h.planService.Get(ctx, req.ID)
}

// Create creates a new plan.
func (h *planHandler) Create(ctx context.Context, req *dto.CreatePlanRequest) (*dto.PlanResponse, error) {
	return h.planService.Create(ctx, req)
}

// Update updates an existing plan.
func (h *planHandler) Update(ctx context.Context, req *dto.UpdatePlanRequest) (*dto.PlanResponse, error) {
	return h.planService.Update(ctx, req.ID, req)
}

// Delete removes a plan by ID.
func (h *planHandler) Delete(ctx context.Context, req *dto.DeletePlanRequest) (*dto.PlanResponse, error) {
	return nil, h.planService.Delete(ctx, req.ID)
}

// FindAll retrieves all plans.
func (h *planHandler) FindAll(ctx context.Context, req *dto.FindPlanRequest) ([]*dto.PlanResponse, error) {
	return h.planService.FindAll(ctx)
}

// FindActive retrieves all active plans.
func (h *planHandler) FindActive(ctx context.Context, req *dto.FindActivePlanRequest) ([]*dto.PlanResponse, error) {
	return h.planService.FindActive(ctx)
}
