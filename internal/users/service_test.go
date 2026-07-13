package users

import (
	"context"
	"errors"
	"testing"
)

type fakeStore struct {
	createFn func(
		ctx context.Context,
		req CreateUserRequest,
	) (User, error)

	createCalls int
}

func (f *fakeStore) Create(
	ctx context.Context,
	req CreateUserRequest,
) (User, error) {
	f.createCalls++

	if f.createFn == nil {
		panic("createFn was not configured")
	}

	return f.createFn(ctx, req)
}

func (f *fakeStore) List(
	ctx context.Context,
) ([]User, error) {
	panic("not implemented")
}

func (f *fakeStore) GetByID(
	ctx context.Context,
	id int,
) (User, error) {
	panic("not implemented")
}

func TestUserServiceCreate(t *testing.T) {
	tests := []struct {
		name           string
		request        CreateUserRequest
		createFn       func(context.Context, CreateUserRequest) (User, error)
		wantErr        error
		wantCreateCall int
	}{
		{
			name: "successful creation",
			request: CreateUserRequest{
				Name:  "Gabriel",
				Email: "gabriel@example.com",
			},
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
			wantCreateCall: 1,
		},
		{
			name: "validation error",
			request: CreateUserRequest{
				Name:  "",
				Email: "gabriel@example.com",
			},
			wantErr:        ValidationError{},
			wantCreateCall: 0,
		},
		{
			name: "store returns error",
			request: CreateUserRequest{
				Name:  "Gabriel",
				Email: "gabriel@example.com",
			},
			createFn: func(
				ctx context.Context,
				req CreateUserRequest,
			) (User, error) {
				return User{}, ErrInternal
			},
			wantErr:        ErrInternal,
			wantCreateCall: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeStore{
				createFn: tt.createFn,
			}

			service := NewService(store)

			user, err := service.Create(
				context.Background(),
				tt.request,
			)

			if tt.wantErr != nil {
				var validationErr ValidationError
				if !errors.Is(err, tt.wantErr) && !errors.As(err, &validationErr) {
					t.Fatalf(
						"expected error %v, got %v",
						tt.wantErr,
						err,
					)
				}
			} else if err != nil {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			if store.createCalls != tt.wantCreateCall {
				t.Fatalf(
					"expected Create() to be called %d times, got %d",
					tt.wantCreateCall,
					store.createCalls,
				)
			}

			if err == nil {
				if user.Name != tt.request.Name {
					t.Fatalf(
						"expected name %q, got %q",
						tt.request.Name,
						user.Name,
					)
				}

				if user.Email != tt.request.Email {
					t.Fatalf(
						"expected email %q, got %q",
						tt.request.Email,
						user.Email,
					)
				}
			}
		})
	}
}
