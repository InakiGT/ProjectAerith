package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce_category/domain"
)

type FindAllCommerceCategoriesCommand struct {
	categoryRepo domain.Repository
}

func NewFindAllCommerceCategories(categoryRepo domain.Repository) *FindAllCommerceCategoriesCommand {
	return &FindAllCommerceCategoriesCommand{
		categoryRepo: categoryRepo,
	}
}

func (cmd *FindAllCommerceCategoriesCommand) Execute(ctx context.Context) ([]*domain.CommerceCategory, error) {
	return cmd.categoryRepo.FindAll(ctx)
}
