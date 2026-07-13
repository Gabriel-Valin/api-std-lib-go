package users

import "context"

type fakeStore struct {
	createFn func(
		ctx context.Context,
		req CreateUserRequest,
	) (User, error)
}

func (f *fakeStore) Create(
	ctx context.Context,
	req CreateUserRequest,
) (User, error) {
	if f.createFn == nil {
		panic("createFn was not configured")
	}

	return f.createFn(
		ctx,
		req,
	)
}

func (f *fakeStore) List(
	ctx context.Context,
) ([]User, error) {
	panic("not implemented")
}

func (f *fakeStore) GetByID(
	ctx context.Context,
	id int,

) (User, error) {
	panic("not implemented")
}
