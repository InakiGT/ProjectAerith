package domain

import "context"

type Repository interface {
	Save(ctx context.Context, product *Product) error
	FindAll(ctx context.Context) ([]*Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	// FindByKeyword(ctx context.Context, keyword string) ([]*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}
