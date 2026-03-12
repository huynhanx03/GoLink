package http

import (
	"context"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
)

// getUserID extracts the user ID from the context (populated by auth middleware).
func getUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(constraints.ContextKeyUserID).(string)
	if !ok || userID == "" {
		return "", apperr.New(response.CodeUnauthorized, "unauthorized", 0, nil)
	}
	return userID, nil
}

// getTenantID is an alias for getUserID, used in webhook contexts where the User ID acts as Tenant ID.
func getTenantID(ctx context.Context) (string, error) {
	return getUserID(ctx)
}
