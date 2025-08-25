package application

import (
	"context"

	"rapi-pedidos/src/internal/client/domain"
)

type DeleteClientCommand struct {
	clientRepo domain.Repository
}

func NewDeleteClient(clientRepo domain.Repository) *DeleteClientCommand {
	return &DeleteClientCommand{
		clientRepo: clientRepo,
	}
}

func (cmd *DeleteClientCommand) Execute(ctx context.Context, id string) error {
	return cmd.clientRepo.Delete(ctx, id)
}
