package domain

import "errors"

var (
	ErrInvalidColor            = errors.New("el color el inválido")
	ErrInvalidType             = errors.New("el tipo de vehículo es inválido")
	ErrInvalidCardID           = errors.New("la tarjeta de circulación es inválida")
	ErrInvalidPlate            = errors.New("la placa de circulación es inválida")
	ErrInvalidDeliveryPersonID = errors.New("el repartidor (ID) es inválido")
)
