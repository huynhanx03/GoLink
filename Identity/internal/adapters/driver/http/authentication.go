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
	OAuthCallback(ctx context.Context, req *dto.OAuthCallbackRequest) (*dto.OAuthCallbackResponse, error)
	OAuthRegister(ctx context.Context, req *dto.OAuthRegisterRequest) (*dto.LoginResponse, error)
	LinkOAuth(ctx context.Context, req *dto.OAuthLinkRequest) (*dto.OAuthLinkResponse, error)
	ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error)
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
	userID, err := h.getUserID(ctx)
	if err != nil {
		return nil, err
	}
	return h.authService.ChangePassword(ctx, userID, req)
}

// RefreshToken handles token refresh.
func (h *authenticationHandler) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	userID, err := h.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return h.authService.RefreshToken(ctx, req, userID)
}

// OAuthCallback handles OAuth provider callback.
func (h *authenticationHandler) OAuthCallback(ctx context.Context, req *dto.OAuthCallbackRequest) (*dto.OAuthCallbackResponse, error) {
	return h.authService.OAuthCallback(ctx, req)
}

// OAuthRegister completes OAuth registration with username and password.
func (h *authenticationHandler) OAuthRegister(ctx context.Context, req *dto.OAuthRegisterRequest) (*dto.LoginResponse, error) {
	return h.authService.OAuthRegister(ctx, req)
}

// LinkOAuth links an OAuth account to the current authenticated user.
func (h *authenticationHandler) LinkOAuth(ctx context.Context, req *dto.OAuthLinkRequest) (*dto.OAuthLinkResponse, error) {
	userID, err := h.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return h.authService.LinkOAuth(ctx, userID, req)
}

// ForgotPassword initiates password reset flow.
func (h *authenticationHandler) ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, error) {
	return h.authService.ForgotPassword(ctx, req)
}

// ResetPassword completes password reset with a reset token.
func (h *authenticationHandler) ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error) {
	return h.authService.ResetPassword(ctx, req)
}

// getUserID extracts and validates the UserID from the context.
func (h *authenticationHandler) getUserID(ctx context.Context) (int, error) {
	userIDVal := ctx.Value(constraints.ContextKeyUserID)
	if userIDVal == nil {
		return 0, apperr.New(response.CodeUnauthorized, "user id not found in context", http.StatusUnauthorized, nil)
	}

	userID, ok := userIDVal.(int)
	if !ok {
		return 0, apperr.New(response.CodeInternalError, "invalid user id type", http.StatusInternalServerError, nil)
	}
	return userID, nil
}
