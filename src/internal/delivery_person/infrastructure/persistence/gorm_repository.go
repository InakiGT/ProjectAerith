package persistence

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type DeliveryPerson struct {
	Id uint `gorm:"primaryKey;autoIncrement;column:id"`
	domain.DeliveryPerson
	gorm.Model
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

func (r *GormRepository) Save(ctx context.Context, deliveryperson *domain.DeliveryPerson) error {
	err := r.db.WithContext(ctx).Save(&DeliveryPerson{DeliveryPerson: *deliveryperson}).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.DeliveryPerson, error) {
	var deliveryPersons []*DeliveryPerson
	err := r.db.WithContext(ctx).Find(&deliveryPersons).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.DeliveryPerson
	for _, deliveryPerson := range deliveryPersons {
		result = append(result, &deliveryPerson.DeliveryPerson)
		result[len(result)-1].Id = deliveryPerson.Id
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.DeliveryPerson, error) {
	var deliveryPerson *DeliveryPerson
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&deliveryPerson).Error
	if err != nil {
		return nil, err
	}

	result := &deliveryPerson.DeliveryPerson
	result.Id = deliveryPerson.Id

	return result, nil
}

func (r *GormRepository) FindByLocation(ctx context.Context, location domain.Location) ([]*domain.DeliveryPerson, error) {
	var deliveryPersons []*DeliveryPerson
	radiusKm := 10

	query := `
		SELECT *,
		(
			6371 * acos(
				cos(radians(?)) * cos(radians(latitude)) *
				cos(radians(longitude) - radians(?)) +
				sin(radians(?)) * sin(radians(latitude))
			)
		) AS distance
		FROM delivery_people
		HAVING distance < ?
		ORDER BY distance;
	`

	err := r.db.WithContext(ctx).
		Raw(query,
			location.Latitude, location.Longitude,
			location.Latitude, radiusKm).
		Scan(&deliveryPersons).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.DeliveryPerson
	for _, dp := range deliveryPersons {
		result = append(result, &dp.DeliveryPerson)
		result[len(result)-1].Id = dp.Id
	}

	return result, nil
}

func (r *GormRepository) FindByPersonalID(ctx context.Context, personalID string) (*domain.DeliveryPerson, error) {
	var deliveryPerson *DeliveryPerson
	err := r.db.WithContext(ctx).Where("personal_id = ?", personalID).First(&deliveryPerson).Error
	if err != nil {
		return nil, err
	}

	result := &deliveryPerson.DeliveryPerson
	result.Id = deliveryPerson.Id

	return result, nil
}

func (r *GormRepository) UpdateCurrentLocation(ctx context.Context, id string, location domain.Location) error {
	var deliveryPersonToUpdate DeliveryPerson

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&deliveryPersonToUpdate).Error
	if err != nil {
		return err
	}

	deliveryPersonToUpdate.CurrentLocation = location
	return r.db.WithContext(ctx).Save(&deliveryPersonToUpdate).Error
}

func (r *GormRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	var deliveryPersonToUpdate DeliveryPerson

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&deliveryPersonToUpdate).Error
	if err != nil {
		return err
	}

	deliveryPersonToUpdate.Status = status
	return r.db.WithContext(ctx).Save(&deliveryPersonToUpdate).Error
}

func (r *GormRepository) Update(ctx context.Context, deliveryPerson *domain.DeliveryPerson) error {
	var deliveryPersonToUpdate DeliveryPerson

	err := r.db.WithContext(ctx).Where("id = ?", deliveryPerson.Id).First(&deliveryPersonToUpdate).Error
	if err != nil {
		return err
	}

	deliveryPersonToUpdate.DeliveryPerson = *deliveryPerson
	return r.db.WithContext(ctx).Save(&deliveryPersonToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&DeliveryPerson{}).Error
}
