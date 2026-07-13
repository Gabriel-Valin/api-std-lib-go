//go:build integration

package users

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupDatabase(t *testing.T) *sql.DB {
	t.Helper()

	ctx := context.Background()

	container, err := postgres.Run(ctx,
		"postgres:17",
		postgres.WithDatabase("products"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %v", err)
		}
	})

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to build connection string: %v", err)
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("failed to ping db: %v", err)
	}

	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}

	return db
}

func TestPostgresStore(t *testing.T) {
	db := setupDatabase(t)
	store := NewPostgresStore(db)
	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		t.Cleanup(func() {
			_, _ = db.ExecContext(ctx, "DELETE FROM users")
		})

		user, err := store.Create(ctx, CreateUserRequest{
			Name:  "Gabriel",
			Email: "gabriel@example.com",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if user.ID == 0 {
			t.Fatal("expected a non-zero ID")
		}

		if user.Name != "Gabriel" || user.Email != "gabriel@example.com" {
			t.Fatalf("unexpected user: %+v", user)
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		t.Cleanup(func() {
			_, _ = db.ExecContext(ctx, "DELETE FROM users")
		})

		created, err := store.Create(ctx, CreateUserRequest{
			Name:  "Gabriel",
			Email: "gabriel@example.com",
		})
		if err != nil {
			t.Fatalf("unexpected error creating fixture: %v", err)
		}

		got, err := store.GetByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != created {
			t.Fatalf("expected %+v, got %+v", created, got)
		}
	})

	t.Run("GetByID_NotFound", func(t *testing.T) {
		_, err := store.GetByID(ctx, 999999)
		if !errors.Is(err, ErrUserNotFound) {
			t.Fatalf("expected ErrUserNotFound, got %v", err)
		}
	})

	t.Run("List", func(t *testing.T) {
		t.Cleanup(func() {
			_, _ = db.ExecContext(ctx, "DELETE FROM users")
		})

		if _, err := store.Create(ctx, CreateUserRequest{Name: "A", Email: "a@example.com"}); err != nil {
			t.Fatalf("unexpected error creating fixture: %v", err)
		}
		if _, err := store.Create(ctx, CreateUserRequest{Name: "B", Email: "b@example.com"}); err != nil {
			t.Fatalf("unexpected error creating fixture: %v", err)
		}

		got, err := store.List(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 2 {
			t.Fatalf("expected 2 users, got %d", len(got))
		}
	})
}
