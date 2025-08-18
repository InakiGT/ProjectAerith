package application

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type FindDeliveryPersonByIDCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewFindDeliveryPersonByID(deliveryPersonRepo domain.Repository) *FindDeliveryPersonByIDCommand {
	return &FindDeliveryPersonByIDCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *FindDeliveryPersonByIDCommand) Execute(ctx context.Context, id string) (*domain.DeliveryPerson, error) {
	return cmd.deliveryPersonRepo.FindByID(ctx, id)
}
