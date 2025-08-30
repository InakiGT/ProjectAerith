package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/commerce_category/domain"
)

type UpdateCommerceCategoryCommand struct {
	categoryRepo domain.Repository
}

func NewUpdateCommerceCategory(categoryRepo domain.Repository) *UpdateCommerceCategoryCommand {
	return &UpdateCommerceCategoryCommand{
		categoryRepo: categoryRepo,
	}
}

func (cmd *UpdateCommerceCategoryCommand) Execute(ctx context.Context, id, name string) (*domain.CommerceCategory, error) {
	category, err := cmd.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("la categoría de comercio no existe")
	}

	if err = cmd.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
