package application

import (
	"context"
	"rapi-pedidos/src/internal/user/domain"
)

type FindAllUsersCommand struct {
	userRepo domain.Repository
}

func NewFindAllUsers(repo domain.Repository) *FindAllUsersCommand {
	return &FindAllUsersCommand{
		userRepo: repo,
	}
}

func (fauc *FindAllUsersCommand) Execute(ctx context.Context) ([]*domain.User, error) {
	users, err := fauc.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
