package users

import (
	"context"
	"testing"
)

func TestCreateUserRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateUserRequest
		wantErr bool
	}{
		{
			name: "valid user",
			request: CreateUserRequest{
				Name:  "Gabriel",
				Email: "gabriel@example.com",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			request: CreateUserRequest{
				Name:  "",
				Email: "gabriel@example.com",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			request: CreateUserRequest{
				Name:  "Gabriel",
				Email: "",
			},
			wantErr: true,
		},
		{
			name: "invalid email",
			request: CreateUserRequest{
				Name:  "Gabriel",
				Email: "invalid-email",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()

			if tt.wantErr && err == nil {
				t.Fatal("expected validation error")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestUserServiceCreate(t *testing.T) {
	store := &fakeStore{
		createFn: func(
			ctx context.Context,
			req CreateUserRequest,
		) (User, error) {
			return User{
				ID:    1,
				Name:  req.Name,
				Email: req.Email,
			}, nil
		},
	}

	service := NewService(store)
	user, err := service.Create(

		context.Background(),
		CreateUserRequest{
			Name:  "Gabriel",
			Email: "gabriel@example.com",
		},
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if user.ID != 1 {
		t.Fatal("unexpected id")
	}

	if user.Name != "Gabriel" {
		t.Fatal("unexpected name")
	}
}
