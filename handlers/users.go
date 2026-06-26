package handlers

import (
	"net/http"

	"github.com/Gabriel-Valin/products-api/internal/users"
)

func Users(w http.ResponseWriter, r *http.Request) {
	_ = users.All
	w.Write([]byte("Users endpoint"))
}
