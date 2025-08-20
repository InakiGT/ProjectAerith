package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type CreateVehicleCommand struct {
	vehicleRepo domain.Repository
}

func NewCreateVehicle(vehicleRepo domain.Repository) *CreateVehicleCommand {
	return &CreateVehicleCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *CreateVehicleCommand) Execute(ctx context.Context, color, vtype, plate, cardid string, deliverypersonid uint) (*domain.Vehicle, error) {
	if cardid == "" {
		return nil, errors.New("el número de tarjeta de circulación es requerido")
	}
	if plate == "" {
		return nil, errors.New("el número de placa es requerido")
	}

	vehicle, _ := cmd.vehicleRepo.FindByCardID(ctx, cardid)
	if vehicle != nil {
		return nil, errors.New("el vehículo con ese número de tarjeta de circulación ya está registrado")
	}

	vehicle, _ = cmd.vehicleRepo.FindByPlate(ctx, plate)
	if vehicle != nil {
		return nil, errors.New("el vehículo con ese número de placa ya está registrado")
	}

	vehicle, err := domain.NewVehicle(color, vtype, plate, cardid, deliverypersonid)
	if err != nil {
		return nil, err
	}

	if err = cmd.vehicleRepo.Save(ctx, vehicle); err != nil {
		return nil, err
	}

	return vehicle, nil
}
