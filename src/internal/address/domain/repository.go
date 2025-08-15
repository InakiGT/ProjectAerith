package domain

import "context"

type Repository interface {
	Save(ctx context.Context, address *Address) error
	FindAll(ctx context.Context) ([]*Address, error)
	FindByID(ctx context.Context, id string) (*Address, error)
	Update(ctx context.Context, address *Address) error
	Delete(ctx context.Context, id string) error
}
