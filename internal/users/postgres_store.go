package users

import (
	"context"
	"database/sql"
	"errors"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) List(ctx context.Context) ([]User, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, name, email
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *PostgresStore) GetByID(ctx context.Context, id int) (User, error) {
	var user User

	err := s.db.QueryRowContext(ctx, `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}

		return User{}, err
	}

	return user, nil
}

func (s *PostgresStore) Create(ctx context.Context, req CreateUserRequest) (User, error) {
	var user User

	err := s.db.QueryRowContext(ctx, `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id, name, email
	`, req.Name, req.Email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
