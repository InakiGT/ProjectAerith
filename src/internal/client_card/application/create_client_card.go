package application

import (
	"context"

	"rapi-pedidos/src/internal/client_card/domain"
)

type CreateClientCardCommand struct {
	cardRepo domain.Repository
}

func NewCreateClientCard(cardRepo domain.Repository) *CreateClientCardCommand {
	return &CreateClientCardCommand{
		cardRepo: cardRepo,
	}
}

func (cmd *CreateClientCardCommand) Execute(ctx context.Context, userid uint, provider, expyear, expmonth, last4, brand, servicecustomerid string) (*domain.ClientCard, error) {
	clientCard, err := domain.NewClientCard(userid, provider, expyear, expmonth, last4, brand, servicecustomerid)
	if err != nil {
		return nil, err
	}

	if err = cmd.cardRepo.Save(ctx, clientCard); err != nil {
		return nil, err
	}

	return clientCard, nil
}
