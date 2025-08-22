package domain

import "context"

type Repository interface {
	Save(ctx context.Context, clientcard *ClientCard) error
	FindAll(ctx context.Context) ([]*ClientCard, error)
	FindByID(ctx context.Context, id string) (*ClientCard, error)
	FindByClientID(ctx context.Context, clientid string) ([]*ClientCard, error)
	Update(ctx context.Context, clientcard *ClientCard) error
	Delete(ctx context.Context, id string) error
}
