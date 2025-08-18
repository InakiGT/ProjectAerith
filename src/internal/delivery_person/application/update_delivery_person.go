package application

import (
	"context"
	"errors"
	"time"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type UpdateDeliveryPersonCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewUpdateDeliveryPerson(deliveryPersonRepo domain.Repository) *UpdateDeliveryPersonCommand {
	return &UpdateDeliveryPersonCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *UpdateDeliveryPersonCommand) Execute(ctx context.Context, id, personalid string, birthday time.Time, mainvehicle uint) (*domain.DeliveryPerson, error) {
	deliveryPerson, err := cmd.deliveryPersonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if deliveryPerson == nil {
		return nil, errors.New("no existe el repartidor a actualizar")
	}

	if err = deliveryPerson.Update(mainvehicle, birthday, personalid); err != nil {
		return nil, err
	}

	if err = cmd.deliveryPersonRepo.Update(ctx, deliveryPerson); err != nil {
		return nil, err
	}

	return deliveryPerson, nil
}
