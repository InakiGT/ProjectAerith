package application

import (
	"context"
	"errors"

	"rapi-pedidos/src/internal/user/domain"
)

type UpdateUserCommand struct {
	userRepo domain.Repository
	hasher   domain.PasswordHasher
}

func NewUpdateUser(repo domain.Repository, hasher domain.PasswordHasher) *UpdateUserCommand {
	return &UpdateUserCommand{
		userRepo: repo,
		hasher:   hasher,
	}
}

func (uuc *UpdateUserCommand) Execute(ctx context.Context, id, name, email, password string) (*domain.User, error) {
	user, _ := uuc.userRepo.FindByID(ctx, id)

	if user == nil {
		return nil, errors.New("el usuario no existe")
	}

	var passwordHased string
	var err error
	if password != "" {
		passwordHased, err = uuc.hasher.Hash(password)
		if err != nil {
			return nil, err
		}
	}

	if err := user.Update(name, email, passwordHased); err != nil {
		return nil, err
	}

	if err := uuc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
