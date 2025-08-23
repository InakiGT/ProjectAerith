package persistence

import (
	"context"
	"time"

	"rapi-pedidos/src/internal/delivery_person/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type DeliveryPerson struct {
	gorm.Model
	UserId          uint
	MainVehicleId   uint
	Birthday        time.Time
	CurrentLocation domain.Location `gorm:"type:point"`
	Status          string
	PersonalID      string `gorm:"unique;not null"`
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

func (r *GormRepository) Save(ctx context.Context, deliveryperson *domain.DeliveryPerson) error {
	gormDeliveryPerson := FromDomainTransformer(deliveryperson)
	err := r.db.WithContext(ctx).Save(gormDeliveryPerson).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.DeliveryPerson, error) {
	var deliveryPersons []*DeliveryPerson
	err := r.db.WithContext(ctx).Find(&deliveryPersons).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.DeliveryPerson
	for _, gormDeliveryPerson := range deliveryPersons {
		deliveryPerson := FromPersistenceTransformer(gormDeliveryPerson)
		result = append(result, deliveryPerson)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.DeliveryPerson, error) {
	var deliveryPerson *DeliveryPerson
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&deliveryPerson).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(deliveryPerson)

	return result, nil
}

func (r *GormRepository) FindByLocation(ctx context.Context, location domain.Location) ([]*domain.DeliveryPerson, error) {
	var deliveryPersons []*DeliveryPerson
	radiusKm := 10

	query := `
		SELECT * FROM delivery_persons;
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
	for _, gormDeliveryPerson := range deliveryPersons {
		deliveryPreson := FromPersistenceTransformer(gormDeliveryPerson)
		result = append(result, deliveryPreson)
	}

	return result, nil
}

func (r *GormRepository) FindByPersonalID(ctx context.Context, personalID string) (*domain.DeliveryPerson, error) {
	var deliveryPerson *DeliveryPerson
	err := r.db.WithContext(ctx).Where("personal_id = ?", personalID).First(&deliveryPerson).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(deliveryPerson)

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

func (r *GormRepository) Update(ctx context.Context, deliveryPerson *domain.DeliveryPerson) error {
	var deliveryPersonToUpdate DeliveryPerson

	err := r.db.WithContext(ctx).Where("id = ?", deliveryPerson.Id).First(&deliveryPersonToUpdate).Error
	if err != nil {
		return err
	}

	deliveryPersonToUpdate = *FromDomainTransformer(deliveryPerson)
	return r.db.WithContext(ctx).Save(&deliveryPersonToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&DeliveryPerson{}).Error
}
