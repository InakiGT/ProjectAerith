package persistence

import (
	"context"
	"time"

	addresspersistence "rapi-pedidos/src/internal/address/infrastructure/persistence"
	"rapi-pedidos/src/internal/commerce/domain"
	categorypersistence "rapi-pedidos/src/internal/commerce_category/infrastructure/persistence"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Commerce struct {
	gorm.Model
	CommerceCategoryId uint
	MainAddressId      uint
	Banner             string
	Status             string
	OpenTime           time.Time
	CloseTime          time.Time
	BaseCommission     float32
	MainAddress        addresspersistence.Address           `gorm:"foreignKey:MainAddressId;references:ID"`
	CommerceCategory   categorypersistence.CommerceCategory `gorm:"foreignKey:CommerceCategoryId;references:ID"`
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, commerce *domain.Commerce) error {
	gormCommerce := FromDomainTransformer(commerce)
	err := r.db.WithContext(ctx).Create(gormCommerce).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Commerce, error) {
	var commerces []*Commerce
	err := r.db.WithContext(ctx).Find(&commerces).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Commerce
	for _, gormCommerce := range commerces {
		commerce := FromPersistenceTransformer(gormCommerce)
		result = append(result, commerce)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Commerce, error) {
	var commerce *Commerce
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&commerce).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(commerce)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, commerce *domain.Commerce) error {
	var commerceToUpdate Commerce

	err := r.db.WithContext(ctx).Where("id = ?", commerce.Id).First(&commerceToUpdate).Error
	if err != nil {
		return err
	}

	commerceToUpdate = *FromDomainTransformer(commerce)
	return r.db.WithContext(ctx).Save(&commerceToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Commerce{}).Error
}
