package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/client/domain"
)

type UpdateClientCommand struct {
	clientRepo domain.Repository
}

func NewUpdateClient(clientRepo domain.Repository) *UpdateClientCommand {
	return &UpdateClientCommand{
		clientRepo: clientRepo,
	}
}

func (cmd *UpdateClientCommand) Execute(ctx context.Context, id string, mainaddressid uint) (*domain.Client, error) {
	client, err := cmd.clientRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, errors.New("no existe el cliente a actualizar")
	}

	if err = client.Update(mainaddressid); err != nil {
		return nil, err
	}

	if err = cmd.clientRepo.Update(ctx, client); err != nil {
		return nil, err
	}

	return client, err
}
