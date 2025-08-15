package domain

import "errors"

var (
	ErrInvalidStreet     = errors.New("Calle inválido")
	ErrInvalidCountry    = errors.New("País inválido")
	ErrInvalidCity       = errors.New("Ciudad inválida")
	ErrInvalidPostalCode = errors.New("Código postal inválido")
	ErrInvalidCologne    = errors.New("Colonia inválida")
	ErrInvalidNumber     = errors.New("Número inválido")
)
