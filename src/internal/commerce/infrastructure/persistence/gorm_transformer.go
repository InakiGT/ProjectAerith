package persistence

import "rapi-pedidos/src/internal/commerce/domain"

func FromDomainTransformer(d *domain.Commerce) *Commerce {
	return &Commerce{
		CommerceCategoryId: d.CommerceCategoryId,
		MainAddressId:      d.MainAddressId,
		Banner:             d.Banner,
		Status:             d.Status,
		OpenTime:           d.OpenTime,
		CloseTime:          d.CloseTime,
		BaseCommission:     d.BaseCommission,
	}
}

func FromPersistenceTransformer(p *Commerce) *domain.Commerce {
	return &domain.Commerce{
		Id:                 p.ID,
		CommerceCategoryId: p.CommerceCategoryId,
		MainAddressId:      p.MainAddressId,
		Banner:             p.Banner,
		Status:             p.Status,
		OpenTime:           p.OpenTime,
		CloseTime:          p.CloseTime,
		BaseCommission:     p.BaseCommission,
	}
}
