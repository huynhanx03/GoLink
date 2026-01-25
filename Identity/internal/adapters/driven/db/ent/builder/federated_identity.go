package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateFederatedIdentity builds the create mutation for FederatedIdentity entity.
func BuildCreateFederatedIdentity(ctx context.Context, e *entity.FederatedIdentity) *generate.FederatedIdentityCreate {
	return global.EntClient.DB(ctx).FederatedIdentity.Create().
		SetUserID(e.UserID).
		SetProvider(e.Provider).
		SetExternalID(e.ExternalID)
}
