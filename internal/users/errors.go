package users

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInternal     = errors.New("internal error")
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
