package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreatePermission builds the create mutation for Permission entity.
func BuildCreatePermission(ctx context.Context, e *entity.Permission) *generate.PermissionCreate {
	builder := global.EntClient.DB(ctx).Permission.Create().
		SetRoleID(e.RoleID).
		SetResourceID(e.ResourceID).
		SetScopes(e.Scopes)
	if e.Description != nil {
		builder.SetDescription(*e.Description)
	}
	return builder
}

// BuildUpdatePermission builds the update mutation for Permission entity.
func BuildUpdatePermission(ctx context.Context, e *entity.Permission) *generate.PermissionUpdateOne {
	builder := global.EntClient.DB(ctx).Permission.UpdateOneID(e.ID).
		SetRoleID(e.RoleID).
		SetResourceID(e.ResourceID).
		SetScopes(e.Scopes)
	if e.Description != nil {
		builder.SetDescription(*e.Description)
	}
	return builder
}
