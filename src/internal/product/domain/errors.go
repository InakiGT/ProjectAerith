package domain

import "errors"

var (
	ErrInvalidName        = errors.New("Nombre de producto inválido")
	ErrInvalidPrice       = errors.New("Precio del producto inválido")
	ErrInvalidDescription = errors.New("Descripción del producto inválida")
)
