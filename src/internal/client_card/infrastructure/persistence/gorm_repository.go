package persistence

import (
	"context"

	"rapi-pedidos/src/internal/client/infrastructure/persistence"
	"rapi-pedidos/src/internal/client_card/domain"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

// Structura específica para GORM, requiere un transformer
type ClientCard struct {
	gorm.Model
	ClientID          uint
	Provider          string
	ExpYear           string
	ExpMonth          string
	Last4             string
	Brand             string
	ServiceCustomerId string
	Client            persistence.Client `gorm:"foreignKey:ClientID;references:ID"`
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(ctx context.Context, clientCard *domain.ClientCard) error {
	gormClientCard := FromDomainTransformer(clientCard)
	err := r.db.WithContext(ctx).Create(gormClientCard).Error

	return err
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*domain.ClientCard, error) {
	var clientCards []*ClientCard
	err := r.db.WithContext(ctx).Find(&clientCards).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.ClientCard
	for _, gormClientCard := range clientCards {
		clientCard := FromPersistenceTransformer(gormClientCard)
		result = append(result, clientCard)
	}

	return result, nil
}

func (r *GormRepository) FindByID(ctx context.Context, id string) (*domain.ClientCard, error) {
	var clientCard *ClientCard
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&clientCard).Error
	if err != nil {
		return nil, err
	}

	result := FromPersistenceTransformer(clientCard)

	return result, nil
}

func (r *GormRepository) FindByClientID(ctx context.Context, clientid string) ([]*domain.ClientCard, error) {
	var clientCards []*ClientCard
	err := r.db.WithContext(ctx).Where("client_id = ?", clientid).Find(&clientCards).Error
	if err != nil {
		return nil, err
	}

	var result []*domain.ClientCard
	for _, gormClientCard := range clientCards {
		clientCard := FromPersistenceTransformer(gormClientCard)
		result = append(result, clientCard)
	}

	return result, nil
}

func (r *GormRepository) Update(ctx context.Context, clientCard *domain.ClientCard) error {
	var cardToUpdate ClientCard

	err := r.db.WithContext(ctx).Where("id = ?", clientCard.Id).First(&cardToUpdate).Error
	if err != nil {
		return err
	}

	cardToUpdate = *FromDomainTransformer(clientCard)
	return r.db.WithContext(ctx).Save(cardToUpdate).Error
}

func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&ClientCard{}).Error
}
