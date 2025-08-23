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
	gorm.Model
	CommerceId  uint ``
	Name        string
	Price       float32
	Description string
	Img         string
	// TODO:
	// Commerce Commerce
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, product *domain.Product) error {
	gormProduct := FromDomainTransformer(product)
	err := r.db.WithContext(ctx).Create(gormProduct).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var products []*Product
	err := r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.Product
	for _, gormProduct := range products {
		product := FromPersistenceTransformer(gormProduct)
		result = append(result, product)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	var product *Product
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(product)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, product *domain.Product) error {
	var productToUpdate Product

	err := r.db.WithContext(ctx).Where("id = ?", product.Id).First(&productToUpdate).Error
	if err != nil {
		return err
	}

	productToUpdate = *FromDomainTransformer(product)
	return r.db.WithContext(ctx).Save(&productToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Product{}).Error
}
