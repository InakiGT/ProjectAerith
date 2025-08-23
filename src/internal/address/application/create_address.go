package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/address/domain"
)

type CreateAddressCommand struct {
	addressRepo domain.Repository
}

func NewCreateAddress(repo domain.Repository) *CreateAddressCommand {
	return &CreateAddressCommand{
		addressRepo: repo,
	}
}

func (cac *CreateAddressCommand) Execute(ctx context.Context, city, country, number, street, postalCode, cologne string, userid uint) (*domain.Address, error) {
	if city == "" || country == "" || number == "" || street == "" || postalCode == "" || cologne == "" {
		return nil, errors.New("todos los campos son requeridos")
	}

	address, err := domain.NewAddress(city, country, number, street, postalCode, cologne, userid)
	if err != nil {
		return nil, err
	}

	err = cac.addressRepo.Save(ctx, address)
	if err != nil {
		return nil, err
	}

	return address, nil
}
