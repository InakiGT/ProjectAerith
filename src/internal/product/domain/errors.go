package domain

import "errors"

var (
	ErrInvalidName        = errors.New("nombre de producto inválido")
	ErrInvalidPrice       = errors.New("precio del producto inválido")
	ErrInvalidDescription = errors.New("descripción del producto inválida")
	ErrInvalidImg         = errors.New("la imagen es inválida")
	ErrInvalidCommerceID  = errors.New("el comercio (ID) es inválido")
)
