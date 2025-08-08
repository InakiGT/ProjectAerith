package persistence

import (
	"database/sql"
	"rapi-pedidos/src/internal/user/domain"
)

type PgRepository struct {
	db *sql.DB
}

func NewPgRepository(db *sql.DB) *PgRepository {
	return &PgRepository{db: db}
}

func (r *PgRepository) Save(user *domain.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)

	return err
}

func (r *PgRepository) FindByID(id string) (*domain.User, error)       { return nil, nil }
func (r *PgRepository) FindByEmail(email string) (*domain.User, error) { return nil, nil }
