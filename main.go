package main

import (
	"log"
	"net/http"

	"github.com/Gabriel-Valin/products-api/handlers"
)

func main() {
	log.Println("Server starting on :8080")

	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/users", handlers.Users)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
