package persistence

import "rapi-pedidos/src/internal/product/domain"

func FromDomainTransformer(d *domain.Product) *Product {
	return &Product{
		CommerceId:  d.CommerceId,
		Name:        d.Name,
		Price:       d.Price,
		Description: d.Description,
		Img:         d.Img,
	}
}

func FromPersistenceTransformer(p *Product) *domain.Product {
	return &domain.Product{
		Id:          p.ID,
		CommerceId:  p.CommerceId,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
		Img:         p.Img,
	}
}
