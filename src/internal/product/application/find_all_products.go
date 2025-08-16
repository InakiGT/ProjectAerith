package application

import (
	"context"
	"rapi-pedidos/src/internal/product/domain"
)

type FindAllProductsCommand struct {
	productRepo domain.Repository
}

func NewFindAllProducts(productRepo domain.Repository) *FindAllProductsCommand {
	return &FindAllProductsCommand{
		productRepo: productRepo,
	}
}

func (fapc *FindAllProductsCommand) Execute(ctx context.Context) ([]*domain.Product, error) {
	return fapc.productRepo.FindAll(ctx)
}
