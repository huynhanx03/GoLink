package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
	"go-link/common/pkg/permissions"

	"github.com/RoaringBitmap/roaring"
)

// RequirePermission checks if the user has the required permission scope for a resource.
func RequirePermission(resourceKey string, requiredScope int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// Admin Bypass
		if isAdmin, ok := ctx.Value(constraints.ContextKeyIsAdmin).(bool); ok && isAdmin {
			c.Next()
			return
		}

		// Permission Check
		perms := ctx.Value(constraints.ContextKeyPermissions)
		if perms == nil {
			response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "permission denied", http.StatusForbidden, nil))
			c.Abort()
			return
		}

		rb, ok := perms.(*roaring.Bitmap)
		if !ok {
			response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "invalid permissions format", http.StatusForbidden, nil))
			c.Abort()
			return
		}

		resourceID := permissions.GetResourceID(resourceKey)
		if resourceID == 0 {
			response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "unknown resource", http.StatusForbidden, nil))
			c.Abort()
			return
		}

		if !permissions.CheckPermission(rb, resourceID, requiredScope) {
			response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "permission denied", http.StatusForbidden, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAdmin checks if the user is an admin.
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if isAdmin, ok := ctx.Value(constraints.ContextKeyIsAdmin).(bool); ok && isAdmin {
			c.Next()
			return
		}

		response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "admin required", http.StatusForbidden, nil))
		c.Abort()
	}
}
