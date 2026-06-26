package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewEncoder(w).Encode(users.All); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
