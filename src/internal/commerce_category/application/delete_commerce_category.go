package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce_category/domain"
)

type DeleteCommerceCategoryCommand struct {
	categoryRepo domain.Repository
}

func NewDeleteCommerceCategory(categoryRepo domain.Repository) *DeleteCommerceCategoryCommand {
	return &DeleteCommerceCategoryCommand{
		categoryRepo: categoryRepo,
	}
}

func (cmd *DeleteCommerceCategoryCommand) Execute(ctx context.Context, id string) error {
	return cmd.categoryRepo.Delete(ctx, id)
}
