package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce_category/domain"
)

type FindCommerceCategoryByIDCommand struct {
	categoryRepo domain.Repository
}

func NewFindCommerceCategoryByID(categoryRepo domain.Repository) *FindCommerceCategoryByIDCommand {
	return &FindCommerceCategoryByIDCommand{
		categoryRepo: categoryRepo,
	}
}

func (cmd *FindCommerceCategoryByIDCommand) Execute(ctx context.Context, id string) (*domain.CommerceCategory, error) {
	return cmd.categoryRepo.FindByID(ctx, id)
}
