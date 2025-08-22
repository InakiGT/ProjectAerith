package application

import (
	"context"

	"rapi-pedidos/src/internal/client_card/domain"
)

type FindAllClientCardsCommand struct {
	cardRepo domain.Repository
}

func NewFindAllClientCards(cardRepo domain.Repository) *FindAllClientCardsCommand {
	return &FindAllClientCardsCommand{
		cardRepo: cardRepo,
	}
}

func (cmd *FindAllClientCardsCommand) Execute(ctx context.Context) ([]*domain.ClientCard, error) {
	return cmd.cardRepo.FindAll(ctx)
}
