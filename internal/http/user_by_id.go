package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func (h *UsersHandler) UserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(ctx, id)
	if err != nil {

		if errors.Is(err, users.ErrUserNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
