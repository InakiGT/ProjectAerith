package domain

import "errors"

var (
	ErrInvalidUserID     = errors.New("referencia al usuario (userid) inválida")
	ErrInvalidVehicleId  = errors.New("referencia al vehículo principal (mainvehicleid) inválida")
	ErrInvalidBirthday   = errors.New("fecha de nacimiento inválida")
	ErrInvalidLocation   = errors.New("localización actual inválida")
	ErrInvalidPersonalID = errors.New("número de identificación oficial inválido")
	ErrInvalidStatus     = errors.New("estado del conductor inválido")
)
