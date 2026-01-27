package builder

import (
	"context"

	"go-link/billing/global"
	"go-link/billing/internal/adapters/driven/db/ent/generate"
	"go-link/billing/internal/adapters/driven/db/ent/generate/subscription"
	"go-link/billing/internal/core/entity"
)

func BuildCreateSubscription(ctx context.Context, e *entity.Subscription) *generate.SubscriptionCreate {
	client := global.EntClient.DB(ctx)
	return client.Subscription.Create().
		SetTenantID(e.TenantID).
		SetPlanID(e.PlanID).
		SetStatus(subscription.Status(e.Status)).
		SetCurrentPeriodStart(e.CurrentPeriodStart).
		SetCurrentPeriodEnd(e.CurrentPeriodEnd).
		SetCancelAtPeriodEnd(e.CancelAtPeriodEnd)
}

func BuildUpdateSubscription(ctx context.Context, e *entity.Subscription) *generate.SubscriptionUpdateOne {
	client := global.EntClient.DB(ctx)
	return client.Subscription.UpdateOneID(e.ID).
		SetPlanID(e.PlanID).
		SetStatus(subscription.Status(e.Status)).
		SetCurrentPeriodEnd(e.CurrentPeriodEnd).
		SetCancelAtPeriodEnd(e.CancelAtPeriodEnd)
}
