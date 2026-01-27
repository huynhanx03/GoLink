package mapper

import (
	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
)

// ToPlanResponse converts Plan entity to PlanResponse DTO.
func ToPlanResponse(e *entity.Plan) *dto.PlanResponse {
	if e == nil {
		return nil
	}
	return &dto.PlanResponse{
		ID:        e.ID,
		Name:      e.Name,
		BasePrice: e.BasePrice,
		Period:    e.Period,
		Limits:    e.Limits,
		IsActive:  e.IsActive,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// ToPlanEntityFromCreate converts CreatePlanRequest to Plan entity.
func ToPlanEntityFromCreate(req *dto.CreatePlanRequest) *entity.Plan {
	return &entity.Plan{
		Name:      req.Name,
		BasePrice: req.BasePrice,
		Period:    req.Period,
		Limits:    req.Limits,
		IsActive:  req.IsActive,
	}
}
