package domain

import "context"

type Repository interface {
	Save(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
