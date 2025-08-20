package application

import (
	"context"

	"rapi-pedidos/src/internal/vehicle/domain"
)

type FindVehicleByCardIDCommand struct {
	vehicleRepo domain.Repository
}

func NewFindVehicleByCardID(vehicleRepo domain.Repository) *FindVehicleByCardIDCommand {
	return &FindVehicleByCardIDCommand{
		vehicleRepo: vehicleRepo,
	}
}

func (cmd *FindVehicleByCardIDCommand) Execute(ctx context.Context, cardid string) (*domain.Vehicle, error) {
	return cmd.vehicleRepo.FindByCardID(ctx, cardid)
}
