package application

import (
	"context"

	"rapi-pedidos/src/internal/client/domain"
)

type FindAllClientsCommand struct {
	clientRepo domain.Repository
}

func NewFindAllClients(clientRepo domain.Repository) *FindAllClientsCommand {
	return &FindAllClientsCommand{
		clientRepo: clientRepo,
	}
}

func (cmd *FindAllClientsCommand) Execute(ctx context.Context) ([]*domain.Client, error) {
	return cmd.clientRepo.FindAll(ctx)
}
