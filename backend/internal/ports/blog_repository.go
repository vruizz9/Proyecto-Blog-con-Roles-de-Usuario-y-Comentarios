package ports

import "blog-backend/internal/domain"

// BlogRepository define las operaciones de persistencia para blogs
type BlogRepository interface {
	Create(blog *domain.Blog) error
	FindByID(id int64) (*domain.Blog, error)
	FindByAuthorID(authorID int64) ([]domain.Blog, error)
	List() ([]domain.Blog, error)
	Update(blog *domain.Blog) error
	Delete(id int64) error
}
