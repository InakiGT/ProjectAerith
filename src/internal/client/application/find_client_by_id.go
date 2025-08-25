package application

import (
	"context"

	"rapi-pedidos/src/internal/client/domain"
)

type FindClientByIDCommand struct {
	clientRepo domain.Repository
}

func NewFindClientByID(clientRepo domain.Repository) *FindClientByIDCommand {
	return &FindClientByIDCommand{
		clientRepo: clientRepo,
	}
}

func (cmd *FindClientByIDCommand) Execute(ctx context.Context, id string) (*domain.Client, error) {
	return cmd.clientRepo.FindByID(ctx, id)
}
