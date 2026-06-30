package users

import "strings"

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
		return ErrNameRequired
	}

	if strings.TrimSpace(r.Email) == "" {
		return ErrEmailRequired
	}

	if !strings.Contains(r.Email, "@") {
		return ErrInvalidEmail
	}

	return nil
}
