package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateUser builds the create mutation for User entity.
func BuildCreateUser(ctx context.Context, e *entity.User) *generate.UserCreate {
	return global.EntClient.DB(ctx).User.Create().
		SetUsername(e.Username).
		SetIsAdmin(e.IsAdmin)
}

// BuildUpdateUser builds the update mutation for User entity.
func BuildUpdateUser(ctx context.Context, e *entity.User) *generate.UserUpdateOne {
	return global.EntClient.DB(ctx).User.UpdateOneID(e.ID).
		SetUsername(e.Username).
		SetIsAdmin(e.IsAdmin)
}
