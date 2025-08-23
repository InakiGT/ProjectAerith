package persistence

import "rapi-pedidos/src/internal/delivery_person/domain"

func FromDomainTransformer(d *domain.DeliveryPerson) *DeliveryPerson {
	return &DeliveryPerson{
		UserId:   d.UserId,
		Birthday: d.Birthday,
		CurrentLocation: domain.Location{
			Latitude:  d.CurrentLocation.Latitude,
			Longitude: d.CurrentLocation.Longitude,
		},
		Status:     d.Status,
		PersonalID: d.PersonalID,
	}
}

func FromPersistenceTransformer(p *DeliveryPerson) *domain.DeliveryPerson {
	return &domain.DeliveryPerson{
		Id:       p.ID,
		UserId:   p.UserId,
		Birthday: p.Birthday,
		CurrentLocation: domain.Location{
			Latitude:  p.CurrentLocation.Latitude,
			Longitude: p.CurrentLocation.Longitude,
		},
		Status:     p.Status,
		PersonalID: p.PersonalID,
	}
}
