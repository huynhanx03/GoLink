package service

import (
	"fmt"
	"go-link/common/pkg/common/apperr"
)

// Generic Action Messages
const (
	MsgCreateFailed  = "failed to create"
	MsgGetFailed     = "failed to get"
	MsgUpdateFailed  = "failed to update"
	MsgDeleteFailed  = "failed to delete"
	MsgCheckFailed   = "failed to check"
	MsgFoundFailed   = "failed to find"
	MsgSaveFailed    = "failed to save"
	MsgGenFailed     = "failed to generate"
	MsgProcessFailed = "failed to process"
	MsgNotFound      = "not found"
)

// Specific Error Messages
const (
	MsgInvalidAuth      = "invalid username or password"
	MsgInvalidCredData  = "invalid credential data"
	MsgUserNoPass       = "user has no password set"
	MsgPassIncorrect    = "current password incorrect"
	MsgUsernameExists   = "username already exists"
	MsgUserNotMember    = "user is not a member of this tenant"
	MsgUnauthorized     = "unauthorized"
	MsgResourceNotFound = "resource not found"
)

// MapError wraps an error with a standardized message"
func MapError(serviceName string, err error, code int, msg string, httpStatus int) *apperr.AppError {
	if err == nil {
		return nil
	}

	formattedMsg := fmt.Sprintf("%s %s", serviceName, msg)
	return apperr.Wrap(err, code, formattedMsg, httpStatus)
}

// NewError creates a new AppError with standardized message format
func NewError(serviceName string, code int, msg string, httpStatus int, cause error) *apperr.AppError {
	formattedMsg := fmt.Sprintf("%s %s", serviceName, msg)
	return apperr.New(code, formattedMsg, httpStatus, cause)
}
