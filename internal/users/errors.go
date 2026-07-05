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

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
