package persistence

import (
	"context"

	"rapi-pedidos/src/internal/delivery_person/infrastructure/persistence"
	"rapi-pedidos/src/internal/vehicle/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Vehicle struct {
	gorm.Model
	DeliveryPersonId uint
	Color            string
	Type             string
	Plate            string
	CardID           string
	DeliveryPerson   persistence.DeliveryPerson `gorm:"foreignKey:DeliveryPersonId"`
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, vehicle *domain.Vehicle) error {
	gormVehicle := FromDomainTransformer(vehicle)
	err := r.db.WithContext(ctx).Create(gormVehicle).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Vehicle, error) {
	var vehicles []*Vehicle
	err := r.db.WithContext(ctx).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Vehicle
	for _, gormVehicle := range vehicles {
		vehicle := FromPersistenceTransformer(gormVehicle)
		result = append(result, vehicle)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Vehicle, error) {
	var vehicle *Vehicle
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(vehicle)

	return result, nil
}

func (r *GormRepository) FindByPlate(ctx context.Context, plate string) (*domain.Vehicle, error) {
	var vehicle *Vehicle
	err := r.db.WithContext(ctx).Where("plate = ?", plate).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(vehicle)

	return result, nil
}

func (r *GormRepository) FindByCardID(ctx context.Context, cardid string) (*domain.Vehicle, error) {
	var vehicle *Vehicle
	err := r.db.WithContext(ctx).Where("card_id = ?", cardid).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(vehicle)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, vehicle *domain.Vehicle) error {
	var vehicleToUpdate Vehicle

	err := r.db.WithContext(ctx).Where("id = ?", vehicle.Id).First(&vehicleToUpdate).Error
	if err != nil {
		return err
	}

	vehicleToUpdate = *FromDomainTransformer(vehicle)
	return r.db.WithContext(ctx).Save(&vehicleToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Vehicle{}).Error
}
