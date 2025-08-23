package application

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type FindDeliveryPersonByLocationCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewFindDeliveryPersonByLocation(deliveryPersonRepo domain.Repository) *FindDeliveryPersonByLocationCommand {
	return &FindDeliveryPersonByLocationCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *FindDeliveryPersonByLocationCommand) Execute(ctx context.Context, location domain.Location) ([]*domain.DeliveryPerson, error) {
	return cmd.deliveryPersonRepo.FindByLocation(ctx, location)
}
