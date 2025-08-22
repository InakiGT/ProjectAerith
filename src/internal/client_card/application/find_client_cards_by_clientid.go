package application

import (
	"context"

	"rapi-pedidos/src/internal/client_card/domain"
)

type FindClientCardsByClientIDCommand struct {
	cardRepo domain.Repository
}

func NewFindClientCardsByClientID(cardRepo domain.Repository) *FindClientCardsByClientIDCommand {
	return &FindClientCardsByClientIDCommand{
		cardRepo: cardRepo,
	}
}

func (cmd *FindClientCardsByClientIDCommand) Execute(ctx context.Context, clientid string) ([]*domain.ClientCard, error) {
	return cmd.cardRepo.FindByClientID(ctx, clientid)
}
