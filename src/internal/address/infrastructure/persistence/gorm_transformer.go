package persistence

import "rapi-pedidos/src/internal/address/domain"

func FromDomainTransformer(d *domain.Address) *Address {
	return &Address{
		City:       d.City,
		Country:    d.Country,
		Number:     d.Number,
		Street:     d.Street,
		PostalCode: d.PostalCode,
		Cologne:    d.Cologne,
		UserID:     d.UserId,
	}
}

func FromPersistenceTransformer(p *Address) *domain.Address {
	return &domain.Address{
		Id:         p.ID,
		City:       p.City,
		Country:    p.Country,
		Number:     p.Number,
		Street:     p.Street,
		PostalCode: p.PostalCode,
		Cologne:    p.Cologne,
		UserId:     p.UserID,
	}
}
