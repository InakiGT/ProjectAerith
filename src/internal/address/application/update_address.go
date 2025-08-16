package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/address/domain"
)

type UpadteAddressCommand struct {
	addressRepo domain.Repository
}

func NewUpdateAddress(addressRepo domain.Repository) *UpadteAddressCommand {
	return &UpadteAddressCommand{
		addressRepo: addressRepo,
	}
}

func (uac *UpadteAddressCommand) Execute(ctx context.Context, id, city, country, number, street, postalCode, cologne string) (*domain.Address, error) {
	address, _ := uac.addressRepo.FindByID(ctx, id)

	if address == nil {
		return nil, errors.New("la dirección no existe")
	}

	if city != "" {
		address.City = city
	}
	if country != "" {
		address.Country = country
	}
	if number != "" {
		address.Number = number
	}
	if street != "" {
		address.Street = street
	}
	if postalCode != "" {
		address.PostalCode = postalCode
	}
	if cologne != "" {
		address.Cologne = cologne
	}

	err := uac.addressRepo.Update(ctx, address)
	if err != nil {
		return nil, err
	}

	return address, nil
}
