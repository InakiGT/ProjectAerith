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
	Id string `gorm:"primaryKey;column:id"`
	domain.User
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, user *domain.User) error {
	err := r.db.WithContext(ctx).Create(&User{User: *user}).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*User
	err := r.db.WithContext(ctx).Find(&users).Error

	if err != nil {
		return nil, err
	}

	var result []*domain.User
	for _, user := range users {
		result = append(result, &user.User)
		result[len(result)-1].Id = user.Id
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user *User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	var result *domain.User
	result = &user.User
	result.Id = user.Id

	return result, nil
}

func (r *GormRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	var result *domain.User
	result = &user.User
	result.Id = user.Id

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, user *domain.User) error {
	var userToUpdate User

	err := r.db.WithContext(ctx).Where("id = ?", user.Id).First(&userToUpdate).Error
	if err != nil {
		return err
	}

	userToUpdate.User = *user
	return r.db.WithContext(ctx).Save(&userToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error
}
