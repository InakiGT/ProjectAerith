package persistence

import "rapi-pedidos/src/internal/client/domain"

func FromDomainTransformer(d *domain.Client) *Client {
	return &Client{
		UserID:        d.UserId,
		MainAddressID: d.MainAddressId,
	}
}

func FromPersistenceTransformer(p *Client) *domain.Client {
	return &domain.Client{
		Id:            p.ID,
		MainAddressId: p.MainAddressID,
		UserId:        p.UserID,
	}
}
