package services

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
)

// UserService implementa los casos de uso para gestión de usuarios
type UserService struct {
	userRepo   ports.UserRepository
	authService ports.AuthService
}

// NewUserService crea una nueva instancia del servicio de usuario
func NewUserService(userRepo ports.UserRepository, authService ports.AuthService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		authService: authService,
	}
}

// Register registra un nuevo usuario
func (s *UserService) Register(username, password string, role domain.Role) (*domain.User, error) {
	// Verificar si el usuario ya existe
	existingUser, _ := s.userRepo.FindByUsername(username)
	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// Hash de la contraseña
	hashedPassword, err := s.authService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Crear nuevo usuario
	user := &domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// No retornar la contraseña hasheada
	user.Password = ""
	return user, nil
}

// GetUserByID obtiene un usuario por su ID
func (s *UserService) GetUserByID(id int64) (*domain.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	user.Password = ""
	return user, nil
}

// GetUserByUsername obtiene un usuario por su nombre de usuario
func (s *UserService) GetUserByUsername(username string) (*domain.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	user.Password = ""
	return user, nil
}

// ListUsers lista todos los usuarios
func (s *UserService) ListUsers() ([]domain.User, error) {
	users, err := s.userRepo.List()
	if err != nil {
		return nil, err
	}
	
	// Ocultar contraseñas
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(id int64, username string, role domain.Role) (*domain.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	user.Username = username
	user.Role = role

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

// DeleteUser elimina un usuario
func (s *UserService) DeleteUser(id int64) error {
	return s.userRepo.Delete(id)
}
