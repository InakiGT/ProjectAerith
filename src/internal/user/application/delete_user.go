package application

import (
	"context"
	"rapi-pedidos/src/internal/user/domain"
)

type DeleteUserCommand struct {
	userRepo domain.Repository
}

func NewDeleteUser(userRepo domain.Repository) *DeleteUserCommand {
	return &DeleteUserCommand{
		userRepo: userRepo,
	}
}

func (duc *DeleteUserCommand) Execute(ctx context.Context, id string) error {
	return duc.userRepo.Delete(ctx, id)
}
