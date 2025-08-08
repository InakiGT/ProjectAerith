package application

import (
	"context"
	"rapi-pedidos/src/internal/user/domain"
)

type FindUserByIdCommand struct {
	userRepo domain.Repository
}

func NewFindUserById(repo domain.Repository) *FindUserByIdCommand {
	return &FindUserByIdCommand{
		userRepo: repo,
	}
}

func (fai *FindUserByIdCommand) Execute(ctx context.Context, id string) (*domain.User, error) {
	user, err := fai.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
