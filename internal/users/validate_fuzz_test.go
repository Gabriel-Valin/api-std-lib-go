package users

import (
	"errors"
	"strings"
	"testing"
)

func FuzzValidate(f *testing.F) {
	f.Add("Gabriel", "gabriel@example.com")
	f.Add("", "gabriel@example.com")
	f.Add("Gabriel", "")
	f.Add("Gabriel", "not-an-email")
	f.Add(" ", " ")
	f.Add("Gabriel", "@")

	f.Fuzz(func(t *testing.T, name, email string) {
		req := CreateUserRequest{Name: name, Email: email}

		err := req.Validate()

		if err == nil {
			if strings.TrimSpace(name) == "" {
				t.Fatalf("accepted empty name %q", name)
			}
			if strings.TrimSpace(email) == "" {
				t.Fatalf("accepted empty email %q", email)
			}
			if !strings.Contains(email, "@") {
				t.Fatalf("accepted email without @ %q", email)
			}
			return
		}

		var validationErr ValidationError
		if !errors.As(err, &validationErr) {
			t.Fatalf("Validate returned a non-ValidationError: %T %v", err, err)
		}
	})
}
