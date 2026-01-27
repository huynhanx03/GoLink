package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateTenantMember builds the create mutation for TenantMember entity.
func BuildCreateTenantMember(ctx context.Context, e *entity.TenantMember) *generate.TenantMemberCreate {
	return global.EntClient.DB(ctx).TenantMember.Create().
		SetTenantID(e.TenantID).
		SetUserID(e.UserID).
		SetRoleID(e.RoleID)
}

// BuildUpdateTenantMember builds the update mutation for TenantMember entity.
func BuildUpdateTenantMember(ctx context.Context, e *entity.TenantMember) *generate.TenantMemberUpdateOne {
	return global.EntClient.DB(ctx).TenantMember.UpdateOneID(e.ID).
		SetTenantID(e.TenantID).
		SetUserID(e.UserID).
		SetRoleID(e.RoleID)
}
