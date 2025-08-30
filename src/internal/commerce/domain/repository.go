package domain

import "context"

type Repository interface {
	Save(ctx context.Context, commerce *Commerce) error
	FindAll(ctx context.Context) ([]*Commerce, error)
	FindByID(ctx context.Context, id string) (*Commerce, error)
	Update(ctx context.Context, commerce *Commerce) error
	Delete(ctx context.Context, id string) error
}
