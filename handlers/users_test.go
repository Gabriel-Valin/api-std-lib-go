package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func TestUsersHandler_Get(t *testing.T) {
	store := users.NewStore()

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
