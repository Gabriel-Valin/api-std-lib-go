package users

import (
	"context"
)

type Service interface {
	List(ctx context.Context) ([]User, error)

	GetByID(ctx context.Context, id int) (User, error)

	Create(ctx context.Context, req CreateUserRequest) (User, error)
}

type UserService struct {
	store Store
}

func NewService(store Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) List(ctx context.Context) ([]User, error) {
	return s.store.List(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id int) (User, error) {
	return s.store.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, req CreateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}

	return s.store.Create(ctx, req)
}
