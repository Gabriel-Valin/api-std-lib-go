package users

import (
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r CreateUserRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return ValidationError{
			Field:   "name",
			Message: "is required",
		}
	}

	if strings.TrimSpace(r.Email) == "" {
		return ValidationError{
			Field:   "email",
			Message: "is required",
		}
	}

	if !strings.Contains(r.Email, "@") {
		return ValidationError{
			Field:   "email",
			Message: "is invalid",
		}
	}

	return nil
}
