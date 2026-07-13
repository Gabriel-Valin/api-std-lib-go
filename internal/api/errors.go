package api

import (
	"errors"
	"net/http"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func statusFromError(err error) int {
	var validationErr users.ValidationError

	switch {
	case errors.As(err, &validationErr):
		return http.StatusBadRequest
	case errors.Is(err, users.ErrUserNotFound):
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}

func messageFromError(err error) string {
	var validationErr users.ValidationError

	switch {
	case errors.As(err, &validationErr):
		return err.Error()
	case errors.Is(err, users.ErrUserNotFound):
		return users.ErrUserNotFound.Error()

	default:
		return "internal server error"
	}
}
