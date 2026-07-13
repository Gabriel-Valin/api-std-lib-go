package users

import (
	"context"
	"testing"
)

func BenchmarkValidate(b *testing.B) {
	cases := map[string]CreateUserRequest{
		"valid":       {Name: "Gabriel", Email: "gabriel@example.com"},
		"empty_name":  {Name: "", Email: "gabriel@example.com"},
		"empty_email": {Name: "Gabriel", Email: ""},
		"bad_email":   {Name: "Gabriel", Email: "gabriel-example.com"},
	}

	for name, req := range cases {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = req.Validate()
			}
		})
	}
}

func BenchmarkUserService_Create(b *testing.B) {
	store := &fakeStore{
		createFn: func(ctx context.Context, req CreateUserRequest) (User, error) {
			return User{ID: 1, Name: req.Name, Email: req.Email}, nil
		},
	}

	service := NewService(store)
	req := CreateUserRequest{Name: "Gabriel", Email: "gabriel@example.com"}
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := service.Create(ctx, req); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
