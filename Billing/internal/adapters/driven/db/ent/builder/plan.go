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

	if e.Features != nil {
		create.SetFeatures(e.Features)
	}

	return create
}

func BuildUpdatePlan(ctx context.Context, e *entity.Plan) *generate.PlanUpdateOne {
	client := global.EntClient.DB(ctx)
	update := client.Plan.UpdateOneID(e.ID).
		SetName(e.Name).
		SetDescription(e.Description).
		SetBasePrice(e.BasePrice).
		SetIsActive(e.IsActive)

	if e.Features != nil {
		update.SetFeatures(e.Features)
	}

	return update
}
