package ports

import (
	"context"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
)

// AuthenticationService defines the authentication business logic interface.
type AuthenticationService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	AcquireToken(ctx context.Context, req *dto.AcquireTokenRequest) (*dto.AcquireTokenResponse, error)
	ChangePassword(ctx context.Context, userID int, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error)
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*entity.User, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int) (*dto.RefreshTokenResponse, error)
}
