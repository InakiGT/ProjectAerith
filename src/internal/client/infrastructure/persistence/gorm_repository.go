package persistence

import (
	"context"

	addresspersistence "rapi-pedidos/src/internal/address/infrastructure/persistence"
	"rapi-pedidos/src/internal/client/domain"
	userpersistence "rapi-pedidos/src/internal/user/infrastructure/persistence"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Client struct {
	gorm.Model
	UserID        uint
	MainAddressID uint
	MainAddress   addresspersistence.Address `gorm:"foreignKey:MainAddressID;references:ID"`
	User          userpersistence.User       `gorm:"foreignKey:UserID;references:ID"`
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, address *domain.Client) error {
	gormAddress := FromDomainTransformer(address)
	err := r.db.WithContext(ctx).Create(gormAddress).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Client, error) {
	var addresses []*Client
	err := r.db.WithContext(ctx).Find(&addresses).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Client
	for _, gormAddress := range addresses {
		address := FromPersistenceTransformer(gormAddress)
		result = append(result, address)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Client, error) {
	var address *Client
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&address).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(address)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, address *domain.Client) error {
	var addressToUpdate Client

	err := r.db.WithContext(ctx).Where("id = ?", address.Id).First(&addressToUpdate).Error
	if err != nil {
		return err
	}

	addressToUpdate = *FromDomainTransformer(address)
	return r.db.WithContext(ctx).Save(&addressToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Client{}).Error
}
