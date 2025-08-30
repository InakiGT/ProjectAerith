package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce/domain"
)

type FindCommerceByIDCommand struct {
	commerceRepo domain.Repository
}

func NewFindCommerceByID(commerceRepo domain.Repository) *FindCommerceByIDCommand {
	return &FindCommerceByIDCommand{
		commerceRepo: commerceRepo,
	}
}

func (cmd *FindCommerceByIDCommand) Execute(ctx context.Context, id string) (*domain.Commerce, error) {
	return cmd.commerceRepo.FindByID(ctx, id)
}
