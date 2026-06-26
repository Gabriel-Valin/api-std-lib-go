package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users.Mu.RLock()
		defer users.Mu.RUnlock()

		if err := json.NewEncoder(w).Encode(users.All); err != nil {
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

		users.Mu.Lock()
		newUser := users.User{
			ID:    strconv.Itoa(len(users.All) + 1),
			Name:  req.Name,
			Email: req.Email,
		}

		users.All = append(users.All, newUser)
		users.Mu.Unlock()

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(req); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
