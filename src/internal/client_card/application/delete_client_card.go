package application

import (
	"context"

	"rapi-pedidos/src/internal/client_card/domain"
)

type DeleteClientCardCommand struct {
	cardRepo domain.Repository
}

func NewDeleteClientCard(cardRepo domain.Repository) *DeleteClientCardCommand {
	return &DeleteClientCardCommand{
		cardRepo: cardRepo,
	}
}

func (cmd *DeleteClientCardCommand) Execute(ctx context.Context, id string) error {
	return cmd.cardRepo.Delete(ctx, id)
}
