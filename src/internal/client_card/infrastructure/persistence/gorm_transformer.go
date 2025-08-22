package persistence

import "rapi-pedidos/src/internal/client_card/domain"

func FromDomainTransformer(d *domain.ClientCard) *ClientCard {
	return &ClientCard{
		ClietdId:          d.ClientId,
		Provider:          d.Provider,
		ExpYear:           d.ExpYear,
		ExpMonth:          d.ExpMonth,
		Last4:             d.Last4,
		Brand:             d.Brand,
		ServiceCustomerId: d.ServiceCustomerId,
	}
}

func FromPersistenceTransformer(p *ClientCard) *domain.ClientCard {
	return &domain.ClientCard{
		Id:                p.ID,
		ClientId:          p.ClietdId,
		Provider:          p.Provider,
		ExpYear:           p.ExpYear,
		ExpMonth:          p.ExpMonth,
		Last4:             p.Last4,
		Brand:             p.Brand,
		ServiceCustomerId: p.ServiceCustomerId,
	}
}
