package application

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type FindDeliveryPersonByPersonalIDCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewFindDeliveryPersonByPersonalID(deliveryPersonRepo domain.Repository) *FindDeliveryPersonByPersonalIDCommand {
	return &FindDeliveryPersonByPersonalIDCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *FindDeliveryPersonByPersonalIDCommand) Execute(ctx context.Context, personalid string) (*domain.DeliveryPerson, error) {
	return cmd.deliveryPersonRepo.FindByPersonalID(ctx, personalid)
}
