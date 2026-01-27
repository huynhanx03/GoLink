package dto

// RegisterRequest represents the payload for user registration.
type RegisterRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required,max=100"`
	LastName  string `json:"last_name" validate:"required,max=100"`
	Gender    int    `json:"gender" validate:"oneof=0 1 2"` // 0: male, 1: female, 2: other
	Birthday  string `json:"birthday" validate:"required,datetime=2006-01-02"`
}

// RegisterResponse represents the response for user registration.
type RegisterResponse struct {
	Success bool `json:"success"`
}

// LoginRequest represents the payload for user login.
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represents the response for user login.
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AcquireTokenRequest represents the payload for acquiring a tenant token.
type AcquireTokenRequest struct {
	TenantID int `json:"tenant_id" validate:"required"`
	UserID   int `json:"-"` // From context
}

// AcquireTokenResponse represents the response for acquiring a tenant token.
type AcquireTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// ChangePasswordRequest represents the payload for changing password.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}

// ChangePasswordResponse represents the response for changing password.
type ChangePasswordResponse struct {
	Success bool `json:"success"`
}

// RefreshTokenRequest represents the payload for refreshing access token.
type RefreshTokenRequest struct {
	TenantID int `json:"tenant_id"`
}

// RefreshTokenResponse represents the response for refreshing access token.
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
