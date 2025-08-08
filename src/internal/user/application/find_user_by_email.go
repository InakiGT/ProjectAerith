package application

import (
	"context"
	"rapi-pedidos/src/internal/user/domain"
)

type FindUserByEmailCommand struct {
	userRepo domain.Repository
}

func NewFindUserByEmail(repo domain.Repository) *FindUserByEmailCommand {
	return &FindUserByEmailCommand{
		userRepo: repo,
	}
}

func (fae *FindUserByEmailCommand) Execute(ctx context.Context, id string) (*domain.User, error) {
	user, err := fae.userRepo.FindByEmail(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
