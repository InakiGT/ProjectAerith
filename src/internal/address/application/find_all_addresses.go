package application

import (
	"context"
	"rapi-pedidos/src/internal/address/domain"
)

type FindAllAddressesCommand struct {
	addressRepo domain.Repository
}

func NewFindAllAddresses(addressRepo domain.Repository) *FindAllAddressesCommand {
	return &FindAllAddressesCommand{
		addressRepo: addressRepo,
	}
}

func (fac *FindAllAddressesCommand) Execute(ctx context.Context) ([]*domain.Address, error) {
	return fac.addressRepo.FindAll(ctx)
}
