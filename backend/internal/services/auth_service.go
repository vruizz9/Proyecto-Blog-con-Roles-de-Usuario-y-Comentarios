package services

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
)

// AuthService implementa los casos de uso para autenticación
type AuthService struct {
	userRepo   ports.UserRepository
	authService ports.AuthService
}

// NewAuthService crea una nueva instancia del servicio de autenticación
func NewAuthService(userRepo ports.UserRepository, authService ports.AuthService) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		authService: authService,
	}
}

// Login autentica un usuario y retorna un token JWT
func (s *AuthService) Login(username, password string) (string, *domain.User, error) {
	// Buscar usuario por username
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, domain.ErrInvalidCredentials
	}

	// Verificar contraseña
	if !s.authService.CheckPassword(password, user.Password) {
		return "", nil, domain.ErrInvalidCredentials
	}

	// Generar token JWT
	token, err := s.authService.GenerateToken(user)
	if err != nil {
		return "", nil, err
	}

	// No retornar la contraseña
	user.Password = ""
	return token, user, nil
}

// ValidateToken valida un token JWT y retorna el usuario
func (s *AuthService) ValidateToken(token string) (*domain.User, error) {
	user, err := s.authService.ValidateToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorized
	}

	// Obtener usuario actualizado de la base de datos
	freshUser, err := s.userRepo.FindByID(user.ID)
	if err != nil {
		return nil, domain.ErrUnauthorized
	}

	freshUser.Password = ""
	return freshUser, nil
}

// ChangePassword cambia la contraseña de un usuario
func (s *AuthService) ChangePassword(userID int64, oldPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return domain.ErrUserNotFound
	}

	// Verificar contraseña actual
	if !s.authService.CheckPassword(oldPassword, user.Password) {
		return domain.ErrInvalidCredentials
	}

	// Hash de la nueva contraseña
	hashedPassword, err := s.authService.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.Update(user)
}
