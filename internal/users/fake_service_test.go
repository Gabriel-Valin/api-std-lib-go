package users

import (
	"context"
)

type fakeService struct {
	listFn func(
		context.Context,
	) ([]User, error)
}

func (f *fakeService) List(
	ctx context.Context,
) ([]User, error) {
	return f.listFn(ctx)
}

func (f *fakeService) GetByID(
	ctx context.Context,
) ([]User, error) {
	panic("not implemented")
}

func (f *fakeService) Create(
	ctx context.Context,
) ([]User, error) {
	panic("not implemented")
}
