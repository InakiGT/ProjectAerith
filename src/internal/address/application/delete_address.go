package application

import (
	"context"
	"rapi-pedidos/src/internal/address/domain"
)

type DeleteAddressCommand struct {
	addressRepo domain.Repository
}

func NewDeleteAddress(addressRepo domain.Repository) *DeleteAddressCommand {
	return &DeleteAddressCommand{
		addressRepo: addressRepo,
	}
}

func (dac *DeleteAddressCommand) Execute(ctx context.Context, id string) error {
	return dac.addressRepo.Delete(ctx, id)
}
