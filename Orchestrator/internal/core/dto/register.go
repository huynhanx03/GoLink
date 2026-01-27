package dto

// RegisterRequest represents the request for user registration.
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
	UserID   int64  `json:"user_id"`
	TenantID int64  `json:"tenant_id"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
}
