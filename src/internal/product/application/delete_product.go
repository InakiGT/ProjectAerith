package application

import (
	"context"

	"rapi-pedidos/src/internal/product/domain"
)

type DeleteProductCommand struct {
	productRepo domain.Repository
}

func NewDeleteProduct(productRepo domain.Repository) *DeleteProductCommand {
	return &DeleteProductCommand{
		productRepo: productRepo,
	}
}

func (dpc *DeleteProductCommand) Execute(ctx context.Context, id string) error {
	return dpc.productRepo.Delete(ctx, id)
}
