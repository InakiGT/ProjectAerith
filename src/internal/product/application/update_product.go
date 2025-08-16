package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/product/domain"
)

type UpdateProductCommand struct {
	productRepo domain.Repository
}

func NewUpdateProduct(productRepo domain.Repository) *UpdateProductCommand {
	return &UpdateProductCommand{
		productRepo: productRepo,
	}
}

func (upc *UpdateProductCommand) Execute(ctx context.Context, id, name, description string, price float32) (*domain.Product, error) {
	product, _ := upc.productRepo.FindByID(ctx, id)

	if product == nil {
		return nil, errors.New("el producto no existe")
	}

	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if price != 0 {
		product.Price = price
	}

	if err := upc.productRepo.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
