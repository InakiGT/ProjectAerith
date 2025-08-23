package persistence

import (
	"context"

	"rapi-pedidos/src/internal/address/domain"
	"rapi-pedidos/src/internal/user/infrastructure/persistence"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Address struct {
	gorm.Model
	City       string
	Country    string
	Number     string
	Street     string
	PostalCode string
	Cologne    string
	UserID     uint
	User       persistence.User `gorm:"foreignKey:UserID;references:ID"`
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, address *domain.Address) error {
	gormAddress := FromDomainTransformer(address)
	err := r.db.WithContext(ctx).Create(gormAddress).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Address, error) {
	var addresses []*Address
	err := r.db.WithContext(ctx).Find(&addresses).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Address
	for _, gormAddress := range addresses {
		address := FromPersistenceTransformer(gormAddress)
		result = append(result, address)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Address, error) {
	var address *Address
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&address).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(address)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, address *domain.Address) error {
	var addressToUpdate Address

	err := r.db.WithContext(ctx).Where("id = ?", address.Id).First(&addressToUpdate).Error
	if err != nil {
		return err
	}

	addressToUpdate = *FromDomainTransformer(address)
	return r.db.WithContext(ctx).Save(&addressToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Address{}).Error
}
