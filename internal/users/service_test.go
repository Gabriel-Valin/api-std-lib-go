package users

import "testing"

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
