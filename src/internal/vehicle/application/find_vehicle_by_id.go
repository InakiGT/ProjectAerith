package application

import (
	"context"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type FindVehicleByIDCommand struct {
	vehicleRepo domain.Repository
}

func NewFindVehicleByID(vehicleRepo domain.Repository) *FindVehicleByIDCommand {
	return &FindVehicleByIDCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *FindVehicleByIDCommand) Execute(ctx context.Context, id string) (*domain.Vehicle, error) {
	return cmd.vehicleRepo.FindByID(ctx, id)
}
