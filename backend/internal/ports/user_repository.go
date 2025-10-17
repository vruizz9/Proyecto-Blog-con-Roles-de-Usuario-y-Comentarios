package ports

import "blog-backend/internal/domain"

// UserRepository define las operaciones de persistencia para usuarios
type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
	FindByID(id int64) (*domain.User, error)
	List() ([]domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
}
