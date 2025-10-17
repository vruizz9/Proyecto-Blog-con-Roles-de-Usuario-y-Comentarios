package ports

import "blog-backend/internal/domain"

// CommentRepository define las operaciones de persistencia para comentarios
type CommentRepository interface {
	Create(comment *domain.Comment) error
	FindByID(id int64) (*domain.Comment, error)
	FindByBlogID(blogID int64) ([]domain.Comment, error)
	FindByUserID(userID int64) ([]domain.Comment, error)
	Update(comment *domain.Comment) error
	Delete(id int64) error
}
