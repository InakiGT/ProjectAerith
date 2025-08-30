package domain

import "errors"

var (
	ErrInvalidMainAddressID      = errors.New("dirección principal (ID) inválida")
	ErrInvalidCommerceCategoryID = errors.New("categoría del comercio inválida")
	ErrInvalidBanner             = errors.New("banner (url) inválido")
	ErrInvalidStatus             = errors.New("status inválido")
	ErrInvalidOpenTime           = errors.New("horario de apertura inválido")
	ErrInvalidCloseTime          = errors.New("horario de cierre inválido")
	ErrInvalidBaseCommission     = errors.New("comisión base inválida")
)
