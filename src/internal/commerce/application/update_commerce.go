package application

import (
	"context"
	"errors"
	"time"

	"rapi-pedidos/src/internal/commerce/domain"
)

type UpdateCommercesCommand struct {
	commerceRepo domain.Repository
}

func NewUpdateCommerce(commerceRepo domain.Repository) *UpdateCommercesCommand {
	return &UpdateCommercesCommand{
		commerceRepo: commerceRepo,
	}
}

func (cmd *UpdateCommercesCommand) Execute(ctx context.Context, id string, mainaddressid, commercecategoryid uint, banner, status string, opentime, closetime time.Time, basecommission float32) (*domain.Commerce, error) {
	commerce, err := cmd.commerceRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if commerce != nil {
		return nil, errors.New("el comercio no existe")
	}

	if err = commerce.Update(mainaddressid, commercecategoryid, banner, status, opentime, closetime, basecommission); err != nil {
		return nil, err
	}

	if err = cmd.commerceRepo.Update(ctx, commerce); err != nil {
		return nil, err
	}

	return commerce, nil
}
