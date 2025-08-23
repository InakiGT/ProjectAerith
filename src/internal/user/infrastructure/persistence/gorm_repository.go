package persistence

import (
	"context"

	"rapi-pedidos/src/internal/user/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, user *domain.User) error {
	gormUser := FromDomainTransformer(user)
	err := r.db.WithContext(ctx).Create(gormUser).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.User
	for _, gormUser := range users {
		user := FromPersistenceTransformer(gormUser)
		result = append(result, user)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user *User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(user)

	return result, nil
}

func (r *GormRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(user)

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, user *domain.User) error {
	var userToUpdate User

	err := r.db.WithContext(ctx).Where("id = ?", user.Id).First(&userToUpdate).Error
	if err != nil {
		return err
	}

	userToUpdate = *FromDomainTransformer(user)
	return r.db.WithContext(ctx).Save(&userToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error
}
