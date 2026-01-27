package builder

import (
	"context"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/core/entity"
)

// BuildCreateResource builds the create mutation for Resource entity.
func BuildCreateResource(ctx context.Context, e *entity.Resource) *generate.ResourceCreate {
	builder := global.EntClient.DB(ctx).Resource.Create().
		SetKey(e.Key)
	if e.Description != nil {
		builder.SetDescription(*e.Description)
	}
	return builder
}

// BuildUpdateResource builds the update mutation for Resource entity.
func BuildUpdateResource(ctx context.Context, e *entity.Resource) *generate.ResourceUpdateOne {
	builder := global.EntClient.DB(ctx).Resource.UpdateOneID(e.ID).
		SetKey(e.Key)
	if e.Description != nil {
		builder.SetDescription(*e.Description)
	}
	return builder
}
