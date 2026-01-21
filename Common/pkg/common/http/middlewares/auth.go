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
	"go-link/common/pkg/permissions"
	"go-link/common/pkg/utils"
)

// Authentication middleware validates the JWT token and sets claims in the context.
func Authentication(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(HeaderAuthorization)
		if authHeader == "" {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "missing authorization header", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != TokenTypeBearer {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "invalid authorization header format", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse token
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "invalid or expired token", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*utils.Claims); ok {
			c.Set(ContextKeyClaims, claims)
			c.Set(ContextKeyUserID, claims.UserID)
			c.Set(ContextKeyUsername, claims.Username)
			c.Set(ContextKeyIsAdmin, claims.IsAdmin)
			c.Set(ContextKeyTenantID, claims.TenantID)
			c.Set(ContextKeyRoles, claims.Roles)

			if claims.PermissionsBlob != "" {
				rb, err := permissions.Decompress(claims.PermissionsBlob)
				if err != nil {
					response.ErrorResponse(c, response.CodeForbidden, apperr.New(response.CodeForbidden, "invalid permission data", http.StatusForbidden, nil))
					c.Abort()
					return
				}
				defer permissions.PutBitmap(rb)
				c.Set(ContextKeyPermissions, rb)
			}

			ctx := context.WithValue(c.Request.Context(), ContextKeyUserID, claims.UserID)
			c.Request = c.Request.WithContext(ctx)
		} else {
			response.ErrorResponse(c, response.CodeUnauthorized, apperr.New(response.CodeUnauthorized, "invalid token claims", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
