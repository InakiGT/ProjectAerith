package application

import (
	"context"

	"rapi-pedidos/src/internal/commerce_category/domain"
)

type CreateCommerceCategoryCommand struct {
	categoryRepo domain.Repository
}

func NewCreateCommerceCategory(categoryRepo domain.Repository) *CreateCommerceCategoryCommand {
	return &CreateCommerceCategoryCommand{
		categoryRepo: categoryRepo,
	}
}

func (cmd *CreateCommerceCategoryCommand) Execute(ctx context.Context, name string) (*domain.CommerceCategory, error) {
	category, err := domain.NewCommerceCategory(name)
	if err != nil {
		return nil, err
	}

	if err = cmd.categoryRepo.Save(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
