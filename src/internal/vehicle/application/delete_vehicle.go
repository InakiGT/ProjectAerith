package application

import (
	"context"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type DeleteVehicleCommand struct {
	vehicleRepo domain.Repository
}

func NewDeleteVehicle(vehicleRepo domain.Repository) *DeleteVehicleCommand {
	return &DeleteVehicleCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *DeleteVehicleCommand) Execute(ctx context.Context, id string) error {
	return cmd.vehicleRepo.Delete(ctx, id)
}
