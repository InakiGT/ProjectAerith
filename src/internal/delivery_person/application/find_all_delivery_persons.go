package application

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type FindAllDeliveryPersonCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewFindAllDeliveryPersons(deliveryPersonRepo domain.Repository) *FindAllDeliveryPersonCommand {
	return &FindAllDeliveryPersonCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *FindAllDeliveryPersonCommand) Execute(ctx context.Context) ([]*domain.DeliveryPerson, error) {
	return cmd.deliveryPersonRepo.FindAll(ctx)
}
