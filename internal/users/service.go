package users

import "context"

type Service interface {
	Create(
		ctx context.Context,
		req CreateUserRequest,
	) (User, error)

	List(
		ctx context.Context,
	) ([]User, error)

	GetByID(
		ctx context.Context,
		id int,
	) (User, error)
}

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) List(
	ctx context.Context,
) ([]User, error) {

	return s.store.List(ctx)
}
