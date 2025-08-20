package domain

import "context"

type Repository interface {
	Save(ctx context.Context, vehicle *Vehicle) error
	FindAll(ctx context.Context) ([]*Vehicle, error)
	FindByID(ctx context.Context, id string) (*Vehicle, error)
	FindByPlate(ctx context.Context, plate string) (*Vehicle, error)
	FindByCardID(ctx context.Context, cardid string) (*Vehicle, error)
	Update(ctx context.Context, vehicle *Vehicle) error
	Delete(ctx context.Context, id string) error
}
