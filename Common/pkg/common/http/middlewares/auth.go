package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
	"go-link/common/pkg/permissions"
	"go-link/common/pkg/utils"
)

// Authentication middleware validates the JWT token and sets claims in the context.
func Authentication(publicKey interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constraints.HeaderAuthorization)
		if authHeader == "" {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "missing authorization header", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != constraints.TokenTypeBearer {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "invalid authorization header format", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse token
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "invalid or expired token", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*utils.Claims); ok {
			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, constraints.ContextKeyClaims, claims)
			ctx = context.WithValue(ctx, constraints.ContextKeyUserID, claims.UserID)
			ctx = context.WithValue(ctx, constraints.ContextKeyUsername, claims.Username)
			ctx = context.WithValue(ctx, constraints.ContextKeyIsAdmin, claims.IsAdmin)
			ctx = context.WithValue(ctx, constraints.ContextKeyTenantID, claims.TenantID)
			ctx = context.WithValue(ctx, constraints.ContextKeyRole, claims.Role)
			ctx = context.WithValue(ctx, constraints.ContextKeyRoleLevel, claims.RoleLevel)
			ctx = context.WithValue(ctx, constraints.ContextKeyTierID, claims.TierID)

			if claims.PermissionsBlob != "" {
				rb, err := permissions.Decompress(claims.PermissionsBlob)
				if err != nil {
					response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "invalid permission data", http.StatusForbidden, nil))
					c.Abort()
					return
				}
				defer permissions.PutBitmap(rb)
				ctx = context.WithValue(ctx, constraints.ContextKeyPermissions, rb)
			}

			c.Request = c.Request.WithContext(ctx)
		} else {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "invalid token claims", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
