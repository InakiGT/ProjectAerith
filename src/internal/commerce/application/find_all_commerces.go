package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce/domain"
)

type FindAllCommercesCommand struct {
	commerceRepo domain.Repository
}

func NewFindAllCommerces(commerceRepo domain.Repository) *FindAllCommercesCommand {
	return &FindAllCommercesCommand{
		commerceRepo: commerceRepo,
	}
}

func (cmd *FindAllCommercesCommand) Execute(ctx context.Context) ([]*domain.Commerce, error) {
	return cmd.commerceRepo.FindAll(ctx)
}
