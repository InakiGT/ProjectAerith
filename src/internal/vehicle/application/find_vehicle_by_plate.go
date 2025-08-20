package application

import (
	"context"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type FindVehicleByPlateCommand struct {
	vehicleRepo domain.Repository
}

func NewFindVehicleByPlate(vehicleRepo domain.Repository) *FindVehicleByPlateCommand {
	return &FindVehicleByPlateCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *FindVehicleByPlateCommand) Execute(ctx context.Context, plate string) (*domain.Vehicle, error) {
	return cmd.vehicleRepo.FindByPlate(ctx, plate)
}
