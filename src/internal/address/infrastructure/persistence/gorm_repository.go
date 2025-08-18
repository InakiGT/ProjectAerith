package persistence

import (
	"context"
	"rapi-pedidos/src/internal/address/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Address struct {
	Id uint `gorm:"primaryKey;autoIncrement;column:id"`
	domain.Address
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, address *domain.Address) error {
	err := r.db.WithContext(ctx).Create(&Address{Address: *address}).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Address, error) {
	var addresses []*Address
	err := r.db.WithContext(ctx).Find(&addresses).Error

	if err != nil {
		return nil, err
	}

	var result []*domain.Address
	for _, address := range addresses {
		result = append(result, &address.Address)
		result[len(result)-1].Id = address.Id
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Address, error) {
	var address *Address
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&address).Error

	if err != nil {
		return nil, err
	}

	var result *domain.Address
	result = &address.Address
	result.Id = address.Id

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, address *domain.Address) error {
	var addressToUpdate Address

	err := r.db.WithContext(ctx).Where("id = ?", address.Id).First(&addressToUpdate).Error
	if err != nil {
		return err
	}

	addressToUpdate.Address = *address
	return r.db.WithContext(ctx).Save(&addressToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Address{}).Error
}
