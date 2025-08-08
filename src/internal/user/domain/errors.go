package domain

import "errors"

var (
	ErrInvalidName     = errors.New("nombre inválido")
	ErrInvalidEmail    = errors.New("email inválido")
	ErrInvalidID       = errors.New("ID inválido")
	ErrInvalidPassword = errors.New("contraseña inválida")
)
