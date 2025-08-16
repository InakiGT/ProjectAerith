package application

import (
	"context"

	"rapi-pedidos/src/internal/product/domain"
)

type FindProductByIdCommand struct {
	productRepo domain.Repository
}

func NewFindProductById(productRepo domain.Repository) *FindProductByIdCommand {
	return &FindProductByIdCommand{
		productRepo: productRepo,
	}
}

func (fpic *FindProductByIdCommand) Execute(ctx context.Context, id string) (*domain.Product, error) {
	return fpic.productRepo.FindByID(ctx, id)
}
