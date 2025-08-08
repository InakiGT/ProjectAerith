package application

import (
	"errors"
	"rapi-pedidos/src/internal/user/domain"
)

type CreateUserCommand struct {
	userRepo domain.Repository
}

func NewCreateUser(repo domain.Repository) *CreateUserCommand {
	return &CreateUserCommand{
		userRepo: repo,
	}
}

func (cuc *CreateUserCommand) Execute(name, email, password string) (*domain.User, error) {
	if email == "" {
		return nil, errors.New("email es requerido")
	}

	exists, _ := cuc.userRepo.FindByEmail(email)
	if exists != nil {
		return nil, errors.New("email ya registrado")
	}

	user, err := domain.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	err = cuc.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
