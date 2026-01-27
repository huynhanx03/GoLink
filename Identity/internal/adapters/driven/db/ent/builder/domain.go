package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateDomain builds the create mutation for Domain entity.
func BuildCreateDomain(ctx context.Context, e *entity.Domain) *generate.DomainCreate {
	builder := global.EntClient.DB(ctx).Domain.Create().
		SetDomain(e.Domain).
		SetTenantID(e.TenantID).
		SetIsVerified(e.IsVerified)

	return builder
}

// BuildUpdateDomain builds the update mutation for Domain entity.
func BuildUpdateDomain(ctx context.Context, e *entity.Domain) *generate.DomainUpdateOne {
	builder := global.EntClient.DB(ctx).Domain.UpdateOneID(e.ID).
		SetDomain(e.Domain).
		SetTenantID(e.TenantID).
		SetIsVerified(e.IsVerified)

	return builder
}
