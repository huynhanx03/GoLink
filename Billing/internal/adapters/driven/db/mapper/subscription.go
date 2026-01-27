package mapper

import (
	"go-link/billing/internal/adapters/driven/db/ent/generate"
	"go-link/billing/internal/core/entity"
)

func ToSubscriptionEntity(e *generate.Subscription) *entity.Subscription {
	if e == nil {
		return nil
	}
	return &entity.Subscription{
		ID:                 e.ID,
		TenantID:           e.TenantID,
		PlanID:             e.PlanID,
		Status:             e.Status.String(),
		CurrentPeriodStart: e.CurrentPeriodStart,
		CurrentPeriodEnd:   e.CurrentPeriodEnd,
		CancelAtPeriodEnd:  e.CancelAtPeriodEnd,
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
	}
}
