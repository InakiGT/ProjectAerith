package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/product/domain"
)

type CreateProductCommand struct {
	productRepo domain.Repository
}

func NewCreateProduct(productRepo domain.Repository) *CreateProductCommand {
	return &CreateProductCommand{
		productRepo: productRepo,
	}
}

func (cpc *CreateProductCommand) Execute(ctx context.Context, name, description string, price float32) (*domain.Product, error) {
	if name == "" || description == "" {
		return nil, errors.New("nombre y descripción del producto son requeridos")
	}

	product, err := domain.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	err = cpc.productRepo.Save(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
