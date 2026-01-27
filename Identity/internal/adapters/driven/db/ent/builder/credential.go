package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateCredential builds the create mutation for Credential entity.
func BuildCreateCredential(ctx context.Context, e *entity.Credential) *generate.CredentialCreate {
	return global.EntClient.DB(ctx).Credential.Create().
		SetUserID(e.UserID).
		SetType(e.Type).
		SetCredentialData(e.CredentialData)
}

// BuildUpdateCredential builds the update mutation for Credential entity.
func BuildUpdateCredential(ctx context.Context, e *entity.Credential) *generate.CredentialUpdateOne {
	return global.EntClient.DB(ctx).Credential.UpdateOneID(e.ID).
		SetUserID(e.UserID).
		SetType(e.Type).
		SetCredentialData(e.CredentialData)
}
