package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenType represents the type of token
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"

	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 7 * 24 * time.Hour
)

// Claims extends standard jwt.Claims
type Claims struct {
	jwt.RegisteredClaims
	UserID    int      `json:"sub_int"`
	Username  string   `json:"username"`
	IsAdmin   bool     `json:"is_admin"`
	TenantID  int      `json:"tid,omitempty"`
	TierID    int      `json:"tier_id,omitempty"`
	RoleLevel int      `json:"role_level,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	PermissionsBlob string    `json:"p,omitempty"`
	Type            TokenType `json:"type"`
}

// GenerateToken generates a JWT token
func GenerateToken(
	secret string,
	userID int,
	username string,
	isAdmin bool,
	tenantID int,
	tierID int,
	roles []string,
	roleLevel int,
	permissionsBlob string,
	tokenType TokenType,
) (string, error) {
	duration := AccessTokenDuration
	if tokenType == RefreshToken {
		duration = RefreshTokenDuration
	}

	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   username,
		},
		UserID:          userID,
		Username:        username,
		IsAdmin:         isAdmin,
		TenantID:        tenantID,
		TierID:          tierID,
		Roles:           roles,
		RoleLevel:       roleLevel,
		PermissionsBlob: permissionsBlob,
		Type:            tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
