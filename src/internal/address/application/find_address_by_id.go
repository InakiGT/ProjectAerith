package application

import (
	"context"
	"rapi-pedidos/src/internal/address/domain"
)

type FindAddresByIdCommand struct {
	addressRepo domain.Repository
}

func NewFindAddressById(addressRepo domain.Repository) *FindAddresByIdCommand {
	return &FindAddresByIdCommand{
		addressRepo: addressRepo,
	}
}

func (fac *FindAddresByIdCommand) Execute(ctx context.Context, id string) (*domain.Address, error) {
	return fac.addressRepo.FindByID(ctx, id)
}
