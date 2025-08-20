package persistence

import (
	"context"

	"rapi-pedidos/src/internal/vehicle/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Vehicle struct {
	Id uint `gorm:"primaryKey;autoIncrement;column:id"`
	domain.Vehicle
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, vehicle *domain.Vehicle) error {
	err := r.db.WithContext(ctx).Create(&Vehicle{Vehicle: *vehicle}).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Vehicle, error) {
	var vehicles []*Vehicle
	err := r.db.WithContext(ctx).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Vehicle
	for _, vehicle := range vehicles {
		result = append(result, &vehicle.Vehicle)
		result[len(result)-1].Id = vehicle.Id
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Vehicle, error) {
	var vehicle *Vehicle
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	result := &vehicle.Vehicle
	result.Id = vehicle.Id

	return result, nil
}

func (r *GormRepository) FindByPlate(ctx context.Context, plate string) (*domain.Vehicle, error) {
	var vehicle *Vehicle
	err := r.db.WithContext(ctx).Where("plate = ?", plate).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	result := &vehicle.Vehicle
	result.Id = vehicle.Id

	return result, nil
}

func (r *GormRepository) FindByCardID(ctx context.Context, cardid string) (*domain.Vehicle, error) {
	var vehicle *Vehicle
	err := r.db.WithContext(ctx).Where("card_id = ?", cardid).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	result := &vehicle.Vehicle
	result.Id = vehicle.Id

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, vehicle *domain.Vehicle) error {
	var vehicleToUpdate Vehicle

	err := r.db.WithContext(ctx).Where("id = ?", vehicle.Id).First(&vehicleToUpdate).Error
	if err != nil {
		return err
	}

	vehicleToUpdate.Vehicle = *vehicle
	return r.db.WithContext(ctx).Save(&vehicleToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Vehicle{}).Error
}
