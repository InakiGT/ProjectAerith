package application

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type DeleteDeliveryPersonCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewDeleteDeliveryPerson(deliveryPersonRepo domain.Repository) *DeleteDeliveryPersonCommand {
	return &DeleteDeliveryPersonCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *DeleteDeliveryPersonCommand) Execute(ctx context.Context, id string) error {
	return cmd.deliveryPersonRepo.Delete(ctx, id)
}
