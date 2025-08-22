package application

import (
	"context"

	"rapi-pedidos/src/internal/client_card/domain"
)

type FindClientCardByIDCommand struct {
	cardRepo domain.Repository
}

func NewFindClientCardByID(cardRepo domain.Repository) *FindClientCardByIDCommand {
	return &FindClientCardByIDCommand{
		cardRepo: cardRepo,
	}
}

func (cmd *FindClientCardByIDCommand) Execute(ctx context.Context, id string) (*domain.ClientCard, error) {
	return cmd.cardRepo.FindByID(ctx, id)
}
