package domain

import "context"

type Repository interface {
	Save(ctx context.Context, commercecategory *CommerceCategory) error
	FindAll(ctx context.Context) ([]*CommerceCategory, error)
	FindByID(ctx context.Context, id string) (*CommerceCategory, error)
	Update(ctx context.Context, commercecategory *CommerceCategory) error
	Delete(ctx context.Context, id string) error
}
