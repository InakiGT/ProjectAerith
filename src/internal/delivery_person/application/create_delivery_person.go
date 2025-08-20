package application

import (
	"context"
	"errors"
	"time"

	"rapi-pedidos/src/internal/delivery_person/domain"
)

type CreateDeliveryPersonCommand struct {
	deliveryPersonRepo domain.Repository
}

func NewCreateDeliveryPerson(deliveryPersonRepo domain.Repository) *CreateDeliveryPersonCommand {
	return &CreateDeliveryPersonCommand{
		deliveryPersonRepo: deliveryPersonRepo,
	}
}

func (cmd *CreateDeliveryPersonCommand) Execute(ctx context.Context, userid uint, birthday time.Time, personalid string) (*domain.DeliveryPerson, error) {
	if personalid == "" {
		return nil, errors.New("el número de identificación oficial es requerido")
	}

	deliveryPerson, err := cmd.deliveryPersonRepo.FindByPersonalID(ctx, personalid)
	if err != nil {
		return nil, err
	}

	if deliveryPerson != nil {
		return nil, errors.New("el repartidor ya está registrado")
	}

	deliveryPerson, err = domain.NewDeliveryPerson(userid, birthday, personalid)
	if err != nil {
		return nil, err
	}

	if err = cmd.deliveryPersonRepo.Save(ctx, deliveryPerson); err != nil {
		return nil, err
	}

	return deliveryPerson, nil
}
