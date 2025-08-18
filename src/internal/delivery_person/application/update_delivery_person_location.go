package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type UpdateDeliveryPersonLocationCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewUpdateDeliveryPersonLocation(deliveryPersonRepo domain.Repository) *UpdateDeliveryPersonLocationCommand {
	return &UpdateDeliveryPersonLocationCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *UpdateDeliveryPersonLocationCommand) Execute(ctx context.Context, id string, location domain.Location) (*domain.DeliveryPerson, error) {
	deliveryPerson, err := cmd.deliveryPersonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if deliveryPerson == nil {
		return nil, errors.New("no existe el repartidor a actualizar")
	}

	if err = deliveryPerson.UpdateCurrentLocation(location); err != nil {
		return nil, err
	}

	if err = cmd.deliveryPersonRepo.UpdateCurrentLocation(ctx, id, location); err != nil {
		return nil, err
	}

	return deliveryPerson, err
}
