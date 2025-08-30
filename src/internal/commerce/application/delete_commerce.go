package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce/domain"
)

type DeleteCommerceCommand struct {
	commerceRepo domain.Repository
}

func NewDeleteCommerce(commerceRepo domain.Repository) *DeleteCommerceCommand {
	return &DeleteCommerceCommand{
		commerceRepo: commerceRepo,
	}
}

func (cmd *DeleteCommerceCommand) Execute(ctx context.Context, id string) error {
	return cmd.commerceRepo.Delete(ctx, id)
}
