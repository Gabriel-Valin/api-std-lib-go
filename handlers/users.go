package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

type UserStore interface {
	List() []users.User
	Create(users.CreateUserRequest) users.User
	GetByID(id string) (users.User, bool)
}

type UsersHandler struct {
	store UserStore
}

func NewUsersHandler(store UserStore) *UsersHandler {
	return &UsersHandler{
		store: store,
	}
}

func (h *UsersHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(h.store.List()); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		var req users.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		user := h.store.Create(req)

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
