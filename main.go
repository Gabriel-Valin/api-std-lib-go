package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting on :8080")

	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
