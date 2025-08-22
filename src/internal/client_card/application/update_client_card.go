package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/client_card/domain"
)

type UpdateClientCardCommand struct {
	cardRepo domain.Repository
}

func NewUpdateClientCard(cardRepo domain.Repository) *UpdateClientCardCommand {
	return &UpdateClientCardCommand{
		cardRepo: cardRepo,
	}
}

func (cmd *UpdateClientCardCommand) Execute(ctx context.Context, id, provider, expyear, expmonth, last4, brand, servicecustomerid string) (*domain.ClientCard, error) {
	clientCard, err := cmd.cardRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if clientCard == nil {
		return nil, errors.New("no se encontró la tarjeta a actualizar")
	}

	if err = clientCard.Update(provider, expyear, expmonth, last4, brand, servicecustomerid); err != nil {
		return nil, err
	}

	if err = cmd.cardRepo.Update(ctx, clientCard); err != nil {
		return nil, err
	}

	return clientCard, nil
}
