package users

import (
	"context"
)

type UserStore interface {
	List(ctx context.Context) ([]User, error)

	Create(
		ctx context.Context,
		req CreateUserRequest,
	) (User, error)

	GetByID(
		ctx context.Context,
		id int,
	) (User, error)
}
