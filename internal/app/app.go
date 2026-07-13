package app

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Gabriel-Valin/products-api/internal/config"
	handlers "github.com/Gabriel-Valin/products-api/internal/http"
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
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(
		"pgx",
		cfg.Database.URL,
	)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)

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
		"/health",
		http.HandlerFunc(handlers.Health),
	)

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
		Addr:         cfg.Server.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
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
