package http

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

type UsersHandler struct {
	service users.Service
}

func NewUsersHandler(service users.Service) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

func (h *UsersHandler) Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {

	case http.MethodGet:

		usersList, err := h.service.List(ctx)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(usersList); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:

		var req users.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		user, err := h.service.Create(ctx, req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
