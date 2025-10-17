package persistence

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
	"database/sql"
	"fmt"
)

// BlogRepositorySQL implementa la interfaz BlogRepository usando SQL
type BlogRepositorySQL struct {
	db *sql.DB
}

// NewBlogRepositorySQL crea una nueva instancia del repositorio SQL de blogs
func NewBlogRepositorySQL(db *sql.DB) ports.BlogRepository {
	return &BlogRepositorySQL{db: db}
}

// Create crea un nuevo blog en la base de datos
func (r *BlogRepositorySQL) Create(blog *domain.Blog) error {
	query := `INSERT INTO blogs (title, content, author_id) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, blog.Title, blog.Content, blog.AuthorID)
	if err != nil {
		return fmt.Errorf("error creando blog: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error obteniendo ID del blog: %w", err)
	}

	blog.ID = id
	return nil
}

// FindByID busca un blog por su ID
func (r *BlogRepositorySQL) FindByID(id int64) (*domain.Blog, error) {
	query := `SELECT id, title, content, author_id FROM blogs WHERE id = ?`
	blog := &domain.Blog{}
	
	err := r.db.QueryRow(query, id).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrBlogNotFound
		}
		return nil, fmt.Errorf("error buscando blog por ID: %w", err)
	}

	return blog, nil
}

// FindByAuthorID busca todos los blogs de un autor
func (r *BlogRepositorySQL) FindByAuthorID(authorID int64) ([]domain.Blog, error) {
	query := `SELECT id, title, content, author_id FROM blogs WHERE author_id = ? ORDER BY id`
	rows, err := r.db.Query(query, authorID)
	if err != nil {
		return nil, fmt.Errorf("error buscando blogs por autor: %w", err)
	}
	defer rows.Close()

	var blogs []domain.Blog
	for rows.Next() {
		var blog domain.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID); err != nil {
			return nil, fmt.Errorf("error escaneando blog: %w", err)
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando blogs: %w", err)
	}

	return blogs, nil
}

// List lista todos los blogs
func (r *BlogRepositorySQL) List() ([]domain.Blog, error) {
	query := `SELECT id, title, content, author_id FROM blogs ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error listando blogs: %w", err)
	}
	defer rows.Close()

	var blogs []domain.Blog
	for rows.Next() {
		var blog domain.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID); err != nil {
			return nil, fmt.Errorf("error escaneando blog: %w", err)
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando blogs: %w", err)
	}

	return blogs, nil
}

// Update actualiza un blog existente
func (r *BlogRepositorySQL) Update(blog *domain.Blog) error {
	query := `UPDATE blogs SET title = ?, content = ? WHERE id = ?`
	result, err := r.db.Exec(query, blog.Title, blog.Content, blog.ID)
	if err != nil {
		return fmt.Errorf("error actualizando blog: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrBlogNotFound
	}

	return nil
}

// Delete elimina un blog por su ID
func (r *BlogRepositorySQL) Delete(id int64) error {
	query := `DELETE FROM blogs WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando blog: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrBlogNotFound
	}

	return nil
}
