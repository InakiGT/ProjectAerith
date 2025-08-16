package persistence

import (
	"context"

	"rapi-pedidos/src/internal/product/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type Product struct {
	Id uint `gorm:"primaryKey;autoIncrement;column:id"`
	domain.Product
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, product *domain.Product) error {
	err := r.db.WithContext(ctx).Create(&Product{Product: *product}).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var products []*Product
	err := r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Product
	for _, product := range products {
		result = append(result, &product.Product)
		result[len(result)-1].Id = product.Id
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	var product *Product
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	result := &product.Product
	result.Id = product.Id

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, product *domain.Product) error {
	var productToUpdate Product

	err := r.db.WithContext(ctx).Where("id = ?", product.Id).First(&productToUpdate).Error
	if err != nil {
		return err
	}

	productToUpdate.Product = *product
	return r.db.WithContext(ctx).Save(&productToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Product{}).Error
}
