package domain

import "errors"

var (
	ErrInvalidStreet     = errors.New("calle inválido")
	ErrInvalidCountry    = errors.New("país inválido")
	ErrInvalidCity       = errors.New("ciudad inválida")
	ErrInvalidPostalCode = errors.New("código postal inválido")
	ErrInvalidCologne    = errors.New("colonia inválida")
	ErrInvalidNumber     = errors.New("número inválido")
)
