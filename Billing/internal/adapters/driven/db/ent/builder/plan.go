package builder

import (
	"context"

	"go-link/billing/global"
	"go-link/billing/internal/adapters/driven/db/ent/generate"
	"go-link/billing/internal/core/entity"
)

func BuildCreatePlan(ctx context.Context, e *entity.Plan) *generate.PlanCreate {
	client := global.EntClient.DB(ctx)
	create := client.Plan.Create().
		SetName(e.Name).
		SetBasePrice(e.BasePrice).
		SetPeriod(e.Period).
		SetIsActive(e.IsActive)

	if e.Limits != nil {
		create.SetLimits(e.Limits)
	}

	return create
}

func BuildUpdatePlan(ctx context.Context, e *entity.Plan) *generate.PlanUpdateOne {
	client := global.EntClient.DB(ctx)
	update := client.Plan.UpdateOneID(e.ID).
		SetName(e.Name).
		SetBasePrice(e.BasePrice).
		SetIsActive(e.IsActive)

	if e.Limits != nil {
		update.SetLimits(e.Limits)
	}

	return update
}
