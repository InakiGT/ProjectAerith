package application

import (
	"context"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type FindAllVehiclesCommand struct {
	vehicleRepo domain.Repository
}

func NewFindAllVehicles(vehicleRepo domain.Repository) *FindAllVehiclesCommand {
	return &FindAllVehiclesCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *FindAllVehiclesCommand) Execute(ctx context.Context) ([]*domain.Vehicle, error) {
	return cmd.vehicleRepo.FindAll(ctx)
}
