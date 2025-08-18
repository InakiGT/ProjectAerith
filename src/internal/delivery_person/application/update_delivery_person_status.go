package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type UpdateDeliveryPersonStatusCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewUpdateDeliveryPersonStatus(deliveryPersonRepo domain.Repository) *UpdateDeliveryPersonStatusCommand {
	return &UpdateDeliveryPersonStatusCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *UpdateDeliveryPersonStatusCommand) Execute(ctx context.Context, id, status string) (*domain.DeliveryPerson, error) {
	deliveryPerson, err := cmd.deliveryPersonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if deliveryPerson == nil {
		return nil, errors.New("no existe el repartidor a actualizar")
	}

	if err = deliveryPerson.UpdateStatus(status); err != nil {
		return nil, err
	}

	if err = cmd.deliveryPersonRepo.UpdateStatus(ctx, id, status); err != nil {
		return nil, err
	}

	return deliveryPerson, nil
}
