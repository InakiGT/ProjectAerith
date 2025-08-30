package persistence

import (
	"context"

	"rapi-pedidos/src/internal/commerce_category/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type CommerceCategory struct {
	gorm.Model
	Name string
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, category *domain.CommerceCategory) error {
	gormCategory := FromDomainTransformer(category)
	err := r.db.WithContext(ctx).Create(gormCategory).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.CommerceCategory, error) {
	var categories []*CommerceCategory
	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.CommerceCategory
	for _, gormCategory := range categories {
		category := FromPersistenceTransformer(gormCategory)
		result = append(result, category)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.CommerceCategory, error) {
	var category *CommerceCategory
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(category)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, category *domain.CommerceCategory) error {
	var categoryToUpdate CommerceCategory

	err := r.db.WithContext(ctx).Where("id = ?", category.Id).First(&categoryToUpdate).Error
	if err != nil {
		return err
	}

	categoryToUpdate = *FromDomainTransformer(category)
	return r.db.WithContext(ctx).Save(&categoryToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&CommerceCategory{}).Error
}
