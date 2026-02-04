package dto

// CreateUserRequest represents the request to create a user.
type CreateUserRequest struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Gender    int
	Birthday  string
}

// CreateUserResponse represents the response after creating a user.
type CreateUserResponse struct {
	UserID int64
}
