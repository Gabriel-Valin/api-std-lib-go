package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

type fakeUserService struct {
	listFn func(context.Context) ([]users.User, error)

	listCalls int
}

func (f *fakeUserService) List(
	ctx context.Context,
) ([]users.User, error) {
	f.listCalls++

	if f.listFn == nil {
		panic("listFn was not configured")
	}

	return f.listFn(ctx)
}

func (f *fakeUserService) Create(
	ctx context.Context,
	req users.CreateUserRequest,
) (users.User, error) {
	panic("Create was not expected to be called")
}

func (f *fakeUserService) GetByID(
	ctx context.Context,
	id int,
) (users.User, error) {
	panic("GetByID was not expected to be called")
}

func TestUsersHandlerList(t *testing.T) {
	expectedUsers := []users.User{
		{
			ID:    1,
			Name:  "Gabriel",
			Email: "gabriel@example.com",
		},
		{
			ID:    2,
			Name:  "Maria",
			Email: "maria@example.com",
		},
	}

	service := &fakeUserService{
		listFn: func(
			ctx context.Context,
		) ([]users.User, error) {
			return expectedUsers, nil
		},
	}

	handler := NewUsersHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.Users(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf(
			"expected Content-Type application/json, got %q",
			contentType,
		)
	}

	var response []users.User

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf(
			"failed to decode response body: %v",
			err,
		)
	}

	if len(response) != len(expectedUsers) {
		t.Fatalf(
			"expected %d users, got %d",
			len(expectedUsers),
			len(response),
		)
	}

	if response[0] != expectedUsers[0] {
		t.Fatalf(
			"expected first user %+v, got %+v",
			expectedUsers[0],
			response[0],
		)
	}

	if service.listCalls != 1 {
		t.Fatalf(
			"expected List to be called once, got %d",
			service.listCalls,
		)
	}
}
