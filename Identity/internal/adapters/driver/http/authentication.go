package http

import (
	"context"
	"net/http"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/ports"
)

// AuthenticationHandler defines the authentication HTTP handler interface.
type AuthenticationHandler interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	AcquireToken(ctx context.Context, req *dto.AcquireTokenRequest) (*dto.AcquireTokenResponse, error)
	ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
}

type authenticationHandler struct {
	handler.BaseHandler
	authService ports.AuthenticationService
}

// NewAuthenticationHandler creates a new AuthenticationHandler instance.
func NewAuthenticationHandler(authService ports.AuthenticationService) AuthenticationHandler {
	return &authenticationHandler{
		authService: authService,
	}
}

// Register handles user registration.
func (h *authenticationHandler) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	return h.authService.Register(ctx, req)
}

// Login handles user login.
func (h *authenticationHandler) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	return h.authService.Login(ctx, req)
}

// AcquireToken handles tenant token acquisition.
func (h *authenticationHandler) AcquireToken(ctx context.Context, req *dto.AcquireTokenRequest) (*dto.AcquireTokenResponse, error) {
	return h.authService.AcquireToken(ctx, req)
}

// ChangePassword handles password change.
func (h *authenticationHandler) ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error) {
	userIDVal := ctx.Value(constraints.ContextKeyUserID)
	if userIDVal == nil {
		return nil, apperr.New(response.CodeUnauthorized, "user id not found in context", http.StatusUnauthorized, nil)
	}

	userID, ok := userIDVal.(int)
	if !ok {
		return nil, apperr.New(response.CodeInternalError, "invalid user id type", http.StatusInternalServerError, nil)
	}
	return h.authService.ChangePassword(ctx, userID, req)
}

// RefreshToken handles token refresh.
func (h *authenticationHandler) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	userIDVal := ctx.Value(constraints.ContextKeyUserID)
	if userIDVal == nil {
		return nil, apperr.New(response.CodeUnauthorized, "user id not found in context", http.StatusUnauthorized, nil)
	}
	userID, ok := userIDVal.(int)
	if !ok {
		return nil, apperr.New(response.CodeInternalError, "invalid user id type", http.StatusInternalServerError, nil)
	}
	
	return h.authService.RefreshToken(ctx, req, userID)
}
