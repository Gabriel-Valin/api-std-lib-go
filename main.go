package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Gabriel-Valin/products-api/handlers"
	"github.com/Gabriel-Valin/products-api/internal/middlewares"
	"github.com/Gabriel-Valin/products-api/internal/users"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open(
		"pgx",
		"postgres://postgres:postgres@localhost:5432/products?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	store := users.NewPostgresStore(db)

	usersHandler := handlers.NewUsersHandler(store)

	http.Handle("/users", middlewares.Logger(http.HandlerFunc(usersHandler.Users)))
	http.HandleFunc("/users/", usersHandler.UserByID)

	log.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
