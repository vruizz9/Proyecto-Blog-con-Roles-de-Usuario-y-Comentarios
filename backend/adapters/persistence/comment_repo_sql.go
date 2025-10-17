package persistence

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
	"database/sql"
	"fmt"
)

// CommentRepositorySQL implementa la interfaz CommentRepository usando SQL
type CommentRepositorySQL struct {
	db *sql.DB
}

// NewCommentRepositorySQL crea una nueva instancia del repositorio SQL de comentarios
func NewCommentRepositorySQL(db *sql.DB) ports.CommentRepository {
	return &CommentRepositorySQL{db: db}
}

// Create crea un nuevo comentario en la base de datos
func (r *CommentRepositorySQL) Create(comment *domain.Comment) error {
	query := `INSERT INTO comments (blog_id, user_id, content) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, comment.BlogID, comment.UserID, comment.Content)
	if err != nil {
		return fmt.Errorf("error creando comentario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error obteniendo ID del comentario: %w", err)
	}

	comment.ID = id
	return nil
}

// FindByID busca un comentario por su ID
func (r *CommentRepositorySQL) FindByID(id int64) (*domain.Comment, error) {
	query := `SELECT id, blog_id, user_id, content FROM comments WHERE id = ?`
	comment := &domain.Comment{}
	
	err := r.db.QueryRow(query, id).Scan(&comment.ID, &comment.BlogID, &comment.UserID, &comment.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrCommentNotFound
		}
		return nil, fmt.Errorf("error buscando comentario por ID: %w", err)
	}

	return comment, nil
}

// FindByBlogID busca todos los comentarios de un blog
func (r *CommentRepositorySQL) FindByBlogID(blogID int64) ([]domain.Comment, error) {
	query := `SELECT id, blog_id, user_id, content FROM comments WHERE blog_id = ? ORDER BY id`
	rows, err := r.db.Query(query, blogID)
	if err != nil {
		return nil, fmt.Errorf("error buscando comentarios por blog: %w", err)
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(&comment.ID, &comment.BlogID, &comment.UserID, &comment.Content); err != nil {
			return nil, fmt.Errorf("error escaneando comentario: %w", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando comentarios: %w", err)
	}

	return comments, nil
}

// FindByUserID busca todos los comentarios de un usuario
func (r *CommentRepositorySQL) FindByUserID(userID int64) ([]domain.Comment, error) {
	query := `SELECT id, blog_id, user_id, content FROM comments WHERE user_id = ? ORDER BY id`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error buscando comentarios por usuario: %w", err)
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(&comment.ID, &comment.BlogID, &comment.UserID, &comment.Content); err != nil {
			return nil, fmt.Errorf("error escaneando comentario: %w", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando comentarios: %w", err)
	}

	return comments, nil
}

// Update actualiza un comentario existente
func (r *CommentRepositorySQL) Update(comment *domain.Comment) error {
	query := `UPDATE comments SET content = ? WHERE id = ?`
	result, err := r.db.Exec(query, comment.Content, comment.ID)
	if err != nil {
		return fmt.Errorf("error actualizando comentario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrCommentNotFound
	}

	return nil
}

// Delete elimina un comentario por su ID
func (r *CommentRepositorySQL) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando comentario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrCommentNotFound
	}

	return nil
}
