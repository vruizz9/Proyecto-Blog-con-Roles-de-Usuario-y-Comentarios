package services

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
)

// BlogService implementa los casos de uso para gestión de blogs
type BlogService struct {
	blogRepo   ports.BlogRepository
	userRepo   ports.UserRepository
}

// NewBlogService crea una nueva instancia del servicio de blog
func NewBlogService(blogRepo ports.BlogRepository, userRepo ports.UserRepository) *BlogService {
	return &BlogService{
		blogRepo: blogRepo,
		userRepo: userRepo,
	}
}

// CreateBlog crea un nuevo blog
func (s *BlogService) CreateBlog(title, content string, authorID int64) (*domain.Blog, error) {
	// Verificar que el autor existe
	_, err := s.userRepo.FindByID(authorID)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	blog := &domain.Blog{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}

	if err := s.blogRepo.Create(blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// GetBlogByID obtiene un blog por su ID
func (s *BlogService) GetBlogByID(id int64) (*domain.Blog, error) {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return nil, domain.ErrBlogNotFound
	}
	return blog, nil
}

// GetBlogsByAuthor obtiene todos los blogs de un autor
func (s *BlogService) GetBlogsByAuthor(authorID int64) ([]domain.Blog, error) {
	// Verificar que el autor existe
	_, err := s.userRepo.FindByID(authorID)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return s.blogRepo.FindByAuthorID(authorID)
}

// ListBlogs lista todos los blogs
func (s *BlogService) ListBlogs() ([]domain.Blog, error) {
	return s.blogRepo.List()
}

// UpdateBlog actualiza un blog existente
func (s *BlogService) UpdateBlog(id int64, title, content string, userID int64, userRole domain.Role) (*domain.Blog, error) {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return nil, domain.ErrBlogNotFound
	}

	// Verificar permisos: solo el autor o un administrador puede editar
	if blog.AuthorID != userID && userRole != domain.RoleAdmin {
		return nil, domain.ErrForbidden
	}

	blog.Title = title
	blog.Content = content

	if err := s.blogRepo.Update(blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// DeleteBlog elimina un blog
func (s *BlogService) DeleteBlog(id int64, userID int64, userRole domain.Role) error {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return domain.ErrBlogNotFound
	}

	// Verificar permisos: solo el autor o un administrador puede eliminar
	if blog.AuthorID != userID && userRole != domain.RoleAdmin {
		return domain.ErrForbidden
	}

	return s.blogRepo.Delete(id)
}
