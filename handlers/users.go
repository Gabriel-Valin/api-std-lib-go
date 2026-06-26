package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := json.NewEncoder(w).Encode(users.List()); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	if r.Method == http.MethodPost {
		var req users.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newUser := users.Create(req)

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(newUser); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
