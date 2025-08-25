package application

import (
	"context"

	"rapi-pedidos/src/internal/client/domain"
)

type CreateClientCommand struct {
	clientRepo domain.Repository
}

func NewCreateClient(clientRepo domain.Repository) *CreateClientCommand {
	return &CreateClientCommand{
		clientRepo: clientRepo,
	}
}

func (cmd *CreateClientCommand) Execute(ctx context.Context, userid, mainaddressid uint) (*domain.Client, error) {
	client, err := domain.NewClient(userid, mainaddressid)
	if err != nil {
		return nil, err
	}

	if err = cmd.clientRepo.Save(ctx, client); err != nil {
		return nil, err
	}

	return client, err
}
