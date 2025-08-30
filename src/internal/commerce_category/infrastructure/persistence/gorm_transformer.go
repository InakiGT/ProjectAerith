package persistence

import "rapi-pedidos/src/internal/commerce_category/domain"

func FromDomainTransformer(d *domain.CommerceCategory) *CommerceCategory {
	return &CommerceCategory{
		Name: d.Name,
	}
}

func FromPersistenceTransformer(p *CommerceCategory) *domain.CommerceCategory {
	return &domain.CommerceCategory{
		Id:   p.ID,
		Name: p.Name,
	}
}
