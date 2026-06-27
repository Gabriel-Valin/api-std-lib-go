package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

type FakeUserStore struct {
	users []users.User
}

func (f *FakeUserStore) List() []users.User {
	return f.users
}

func (f *FakeUserStore) Create(req users.CreateUserRequest) users.User {
	user := users.User{
		ID:    "1",
		Name:  req.Name,
		Email: req.Email,
	}

	f.users = append(f.users, user)

	return user
}

func (f *FakeUserStore) GetByID(id string) (users.User, bool) {
	for _, user := range f.users {
		if user.ID == id {
			return user, true
		}
	}

	return users.User{}, false
}

func TestUsersHandler_Get(t *testing.T) {
	store := &FakeUserStore{

		users: []users.User{
			{
				ID:    "1",
				Name:  "Gabriel",
				Email: "gabriel@email.com",
			},
		},
	}

	handler := NewUsersHandler(store)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)
	rec := httptest.NewRecorder()

	handler.Users(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestUsersHandler_GetByID(t *testing.T) {
	store := &FakeUserStore{

		users: []users.User{
			{
				ID:    "43",
				Name:  "XYZ",
				Email: "gabriel@email.com",
			},
		},
	}

	handler := NewUsersHandler(store)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users/43",
		nil,
	)
	rec := httptest.NewRecorder()

	handler.Users(rec, req)

	body := rec.Body.String()
	t.Log(body)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}
