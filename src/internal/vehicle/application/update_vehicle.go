package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type UpdateVehicleCommand struct {
	vehicleRepo domain.Repository
}

func NewUpdateVehicle(vehicleRepo domain.Repository) *UpdateVehicleCommand {
	return &UpdateVehicleCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *UpdateVehicleCommand) Execute(ctx context.Context, id, color, vtype, plate, cardid string, deliverypersonid uint) error {
	vehicle, err := cmd.vehicleRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if vehicle == nil {
		return errors.New("el vehículo a actualizar no existe")
	}

	if err = vehicle.Update(color, vtype, plate, cardid, deliverypersonid); err != nil {
		return err
	}

	if err = cmd.vehicleRepo.Update(ctx, vehicle); err != nil {
		return err
	}

	return nil
}
