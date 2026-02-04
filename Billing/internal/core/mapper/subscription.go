package mapper

import (
	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
)

// ToSubscriptionResponse converts Subscription entity to SubscriptionResponse DTO.
func ToSubscriptionResponse(e *entity.Subscription) *dto.SubscriptionResponse {
	if e == nil {
		return nil
	}
	return &dto.SubscriptionResponse{
		ID:                 e.ID,
		TenantID:           e.TenantID,
		PlanID:             e.PlanID,
		Status:             e.Status,
		CurrentPeriodStart: e.CurrentPeriodStart,
		CurrentPeriodEnd:   e.CurrentPeriodEnd,
		CancelAtPeriodEnd:  e.CancelAtPeriodEnd,
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
	}
}

// ToSubscriptionEntityFromCreate converts CreateSubscriptionRequest to Subscription entity.
func ToSubscriptionEntityFromCreate(req *dto.CreateSubscriptionRequest) *entity.Subscription {
	return &entity.Subscription{
		TenantID: req.TenantID,
		PlanID:   req.PlanID,
	}
}
