package domain

import "context"

type Repository interface {
	Save(ctx context.Context, deliveryPerson *DeliveryPerson) error
	FindAll(ctx context.Context) (*[]DeliveryPerson, error)
	FindByID(ctx context.Context, id string) (*DeliveryPerson, error)
	FindByLocation(ctx context.Context, location Location) (*[]DeliveryPerson, error)
	FindByPersonalID(ctx context.Context, personalID string) (*DeliveryPerson, error)
	UpdateCurrentLocation(ctx context.Context, id string, location Location) error
	UpdateStatus(ctx context.Context, id string, status string) error
	Update(ctx context.Context, deliveryPerson *DeliveryPerson) error
	Delete(ctx context.Context, id string) error
}
