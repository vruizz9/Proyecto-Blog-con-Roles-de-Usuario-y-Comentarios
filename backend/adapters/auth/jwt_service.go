package auth

import (
	"blog-backend/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// JWTService implementa la interfaz AuthService del dominio
type JWTService struct {
	secretKey []byte
}

// Claims define la estructura del token JWT
type Claims struct {
	UserID   int64        `json:"user_id"`
	Username string       `json:"username"`
	Role     domain.Role  `json:"role"`
	jwt.RegisteredClaims
}

// NewJWTService crea una nueva instancia del servicio JWT
func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
	}
}

// GenerateToken genera un token JWT para un usuario
func (j *JWTService) GenerateToken(user *domain.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// ValidateToken valida un token JWT y retorna el usuario
func (j *JWTService) ValidateToken(tokenString string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inesperado")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &domain.User{
			ID:       claims.UserID,
			Username: claims.Username,
			Role:     claims.Role,
		}, nil
	}

	return nil, errors.New("token inválido")
}

// HashPassword hashea una contraseña usando bcrypt
func (j *JWTService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifica si una contraseña coincide con su hash
func (j *JWTService) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
