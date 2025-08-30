package application

import (
	"context"
	"time"

	"rapi-pedidos/src/internal/commerce/domain"
)

type CreateCommercesCommand struct {
	commerceRepo domain.Repository
}

func NewCreateCommerce(commerceRepo domain.Repository) *CreateCommercesCommand {
	return &CreateCommercesCommand{
		commerceRepo: commerceRepo,
	}
}

func (cmd *CreateCommercesCommand) Execute(ctx context.Context, mainaddressid, commercecategoryid uint, banner, status string, opentime, closetime time.Time, basecommission float32) (*domain.Commerce, error) {
	commerce, err := domain.NewCommerce(mainaddressid, commercecategoryid, banner, status, opentime, closetime, basecommission)
	if err != nil {
		return nil, err
	}

	if err = cmd.commerceRepo.Save(ctx, commerce); err != nil {
		return nil, err
	}

	return commerce, nil
}
