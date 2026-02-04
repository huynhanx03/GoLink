package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateTenant builds the create mutation for Tenant entity.
func BuildCreateTenant(ctx context.Context, e *entity.Tenant) *generate.TenantCreate {
	return global.EntClient.DB(ctx).Tenant.Create().
		SetName(e.Name).
		SetPlanID(e.PlanID)
}

// BuildUpdateTenant builds the update mutation for Tenant entity.
func BuildUpdateTenant(ctx context.Context, e *entity.Tenant) *generate.TenantUpdateOne {
	return global.EntClient.DB(ctx).Tenant.UpdateOneID(e.ID).
		SetName(e.Name).
		SetPlanID(e.PlanID)
}
