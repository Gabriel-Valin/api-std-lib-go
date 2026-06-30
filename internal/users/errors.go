package users

import "errors"

var (
	// Domain errors
	ErrUserNotFound = errors.New("user not found")

	// Validation errors
	ErrNameRequired  = errors.New("name is required")
	ErrEmailRequired = errors.New("email is required")
	ErrInvalidEmail  = errors.New("invalid email")
)
