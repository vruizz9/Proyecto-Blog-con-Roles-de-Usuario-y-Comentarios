package ports

import "blog-backend/internal/domain"

// AuthService define las operaciones de autenticación
type AuthService interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.User, error)
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
}
