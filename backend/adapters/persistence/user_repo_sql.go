package persistence

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
	"database/sql"
	"fmt"
)

// UserRepositorySQL implementa la interfaz UserRepository usando SQL
type UserRepositorySQL struct {
	db *sql.DB
}

// NewUserRepositorySQL crea una nueva instancia del repositorio SQL de usuarios
func NewUserRepositorySQL(db *sql.DB) ports.UserRepository {
	return &UserRepositorySQL{db: db}
}

// Create crea un nuevo usuario en la base de datos
func (r *UserRepositorySQL) Create(user *domain.User) error {
	query := `INSERT INTO users (username, password, role) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("error creando usuario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error obteniendo ID del usuario: %w", err)
	}

	user.ID = id
	return nil
}

// FindByUsername busca un usuario por su nombre de usuario
func (r *UserRepositorySQL) FindByUsername(username string) (*domain.User, error) {
	query := `SELECT id, username, password, role FROM users WHERE username = ?`
	user := &domain.User{}
	
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("error buscando usuario por username: %w", err)
	}

	return user, nil
}

// FindByID busca un usuario por su ID
func (r *UserRepositorySQL) FindByID(id int64) (*domain.User, error) {
	query := `SELECT id, username, password, role FROM users WHERE id = ?`
	user := &domain.User{}
	
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("error buscando usuario por ID: %w", err)
	}

	return user, nil
}

// List lista todos los usuarios
func (r *UserRepositorySQL) List() ([]domain.User, error) {
	query := `SELECT id, username, password, role FROM users ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error listando usuarios: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, fmt.Errorf("error escaneando usuario: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando usuarios: %w", err)
	}

	return users, nil
}

// Update actualiza un usuario existente
func (r *UserRepositorySQL) Update(user *domain.User) error {
	query := `UPDATE users SET username = ?, password = ?, role = ? WHERE id = ?`
	result, err := r.db.Exec(query, user.Username, user.Password, user.Role, user.ID)
	if err != nil {
		return fmt.Errorf("error actualizando usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// Delete elimina un usuario por su ID
func (r *UserRepositorySQL) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
