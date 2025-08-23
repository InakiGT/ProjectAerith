package persistence

import "rapi-pedidos/src/internal/vehicle/domain"

func FromDomainTransformer(d *domain.Vehicle) *Vehicle {
	return &Vehicle{
		DeliveryPersonId: d.DeliveryPersonId,
		Color:            d.Color,
		Type:             d.Type,
		Plate:            d.Plate,
		CardID:           d.CardID,
	}
}

func FromPersistenceTransformer(d *Vehicle) *domain.Vehicle {
	return &domain.Vehicle{
		Id:               d.ID,
		DeliveryPersonId: d.DeliveryPersonId,
		Color:            d.Color,
		Type:             d.Type,
		Plate:            d.Plate,
		CardID:           d.CardID,
	}
}
