package services

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
)

// CommentService implementa los casos de uso para gestión de comentarios
type CommentService struct {
	commentRepo ports.CommentRepository
	blogRepo    ports.BlogRepository
	userRepo    ports.UserRepository
}

// NewCommentService crea una nueva instancia del servicio de comentarios
func NewCommentService(commentRepo ports.CommentRepository, blogRepo ports.BlogRepository, userRepo ports.UserRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		blogRepo:    blogRepo,
		userRepo:    userRepo,
	}
}

// CreateComment crea un nuevo comentario
func (s *CommentService) CreateComment(blogID, userID int64, content string) (*domain.Comment, error) {
	// Verificar que el blog existe
	_, err := s.blogRepo.FindByID(blogID)
	if err != nil {
		return nil, domain.ErrBlogNotFound
	}

	// Verificar que el usuario existe
	_, err = s.userRepo.FindByID(userID)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	comment := &domain.Comment{
		BlogID:  blogID,
		UserID:  userID,
		Content: content,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentByID obtiene un comentario por su ID
func (s *CommentService) GetCommentByID(id int64) (*domain.Comment, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return nil, domain.ErrCommentNotFound
	}
	return comment, nil
}

// GetCommentsByBlog obtiene todos los comentarios de un blog
func (s *CommentService) GetCommentsByBlog(blogID int64) ([]domain.Comment, error) {
	// Verificar que el blog existe
	_, err := s.blogRepo.FindByID(blogID)
	if err != nil {
		return nil, domain.ErrBlogNotFound
	}

	return s.commentRepo.FindByBlogID(blogID)
}

// GetCommentsByUser obtiene todos los comentarios de un usuario
func (s *CommentService) GetCommentsByUser(userID int64) ([]domain.Comment, error) {
	// Verificar que el usuario existe
	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return s.commentRepo.FindByUserID(userID)
}

// UpdateComment actualiza un comentario existente
func (s *CommentService) UpdateComment(id int64, content string, userID int64, userRole domain.Role) (*domain.Comment, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return nil, domain.ErrCommentNotFound
	}

	// Verificar permisos: solo el autor o un administrador puede editar
	if comment.UserID != userID && userRole != domain.RoleAdmin {
		return nil, domain.ErrForbidden
	}

	comment.Content = content

	if err := s.commentRepo.Update(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// DeleteComment elimina un comentario
func (s *CommentService) DeleteComment(id int64, userID int64, userRole domain.Role) error {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return domain.ErrCommentNotFound
	}

	// Verificar permisos: solo el autor o un administrador puede eliminar
	if comment.UserID != userID && userRole != domain.RoleAdmin {
		return domain.ErrForbidden
	}

	return s.commentRepo.Delete(id)
}
