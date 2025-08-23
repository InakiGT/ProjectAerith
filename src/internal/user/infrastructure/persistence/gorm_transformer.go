package persistence

import "rapi-pedidos/src/internal/user/domain"

func FromDomainTransformer(d *domain.User) *User {
	return &User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}

func FromPersistenceTransformer(p *User) *domain.User {
	return &domain.User{
		Id:        p.ID,
		CreatedAt: p.CreatedAt,
		DeletedAt: p.DeletedAt.Time,
		UpdatedAt: p.UpdatedAt,
		Name:      p.Name,
		Email:     p.Email,
		Password:  p.Email,
	}
}
