package main

import (
	"log"
	"net/http"

	"github.com/Gabriel-Valin/products-api/handlers"
	"github.com/Gabriel-Valin/products-api/internal/users"
)

func main() {
	store := users.NewStore()

	usersHandler := handlers.NewUsersHandler(store)

	http.HandleFunc("/users", usersHandler.Users)
	http.HandleFunc("/users/", usersHandler.UserByID)

	log.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
