package domain

import "context"

type Repository interface {
	Save(ctx context.Context, client *Client) error
	FindAll(ctx context.Context) ([]*Client, error)
	FindByID(ctx context.Context, id string) (*Client, error)
	Update(ctx context.Context, client *Client) error
	Delete(ctx context.Context, id string) error
}
