package domain

import "errors"

var (
	ErrInvalidUserID        = errors.New("usuario (ID) inválido")
	ErrInvalidMainAddressID = errors.New("dirección principal (ID) inválida")
)
