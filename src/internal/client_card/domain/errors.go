package domain

import "errors"

var (
	ErrInvalidUserID            = errors.New("el cliente (UserID) es inválido")
	ErrInvalidProvider          = errors.New("el proveedor es inválido")
	ErrInvalidExpYear           = errors.New("el año de expiración es inválido")
	ErrInvalidExpMonth          = errors.New("el mes de expiración es inválido")
	ErrInvalidLast4             = errors.New("el número de tarjeta es inválido")
	ErrInvalidBrand             = errors.New("el banco/marca de la tarjeta es inválido")
	ErrInvalidServiceCustomerID = errors.New("el id del proveedor de servicio de tarjetas es inválido")
)
