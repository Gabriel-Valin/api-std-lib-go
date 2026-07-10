package app

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Gabriel-Valin/products-api/handlers"
	"github.com/Gabriel-Valin/products-api/internal/middlewares"
	"github.com/Gabriel-Valin/products-api/internal/users"
)

type App struct {
	server *http.Server
	db     *sql.DB
}

func New() (*App, error) {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	)

	slog.SetDefault(logger)

	db, err := sql.Open(
		"pgx",
		"postgres://postgres:postgres@localhost:5432/products?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	store := users.NewPostgresStore(db)

	service := users.NewService(store)
	usersHandler := handlers.NewUsersHandler(service)

	mux := http.NewServeMux()

	mux.Handle(
		"/users",
		http.HandlerFunc(usersHandler.Users),
	)

	mux.Handle(
		"/users/",
		http.HandlerFunc(usersHandler.UserByID),
	)

	handler := middlewares.Chain(
		mux,
		middlewares.Recovery,
		middlewares.RequestID,
		middlewares.Logger,
		middlewares.Timing,
	)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &App{
		server: server,
		db:     db,
	}, nil
}

func (a *App) Run() error {
	slog.Info("server starting", "addr", a.server.Addr)

	return a.server.ListenAndServe()
}
