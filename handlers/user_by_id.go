package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *UsersHandler) UserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/users/")

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	user, ok := h.store.GetByID(id)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
