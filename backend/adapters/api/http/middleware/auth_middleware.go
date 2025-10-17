package middleware

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/ports"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica la autenticación del usuario mediante JWT
type AuthMiddleware struct {
	authService ports.AuthService
}

// NewAuthMiddleware crea una nueva instancia del middleware de autenticación
func NewAuthMiddleware(authService ports.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// Authenticate verifica el token JWT y establece el usuario en el contexto
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización requerido"})
			c.Abort()
			return
		}

		// Extraer el token del header "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		user, err := m.authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Establecer el usuario en el contexto
		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("user_role", user.Role)

		c.Next()
	}
}

// RequireRole verifica que el usuario tenga un rol específico
func (m *AuthMiddleware) RequireRole(requiredRole domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
			c.Abort()
			return
		}

		role, ok := userRole.(domain.Role)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
			c.Abort()
			return
		}

		if role != requiredRole && role != domain.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso prohibido"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAdmin verifica que el usuario sea administrador
func (m *AuthMiddleware) RequireAdmin() gin.HandlerFunc {
	return m.RequireRole(domain.RoleAdmin)
}

// OptionalAuth permite acceso opcional con autenticación
func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Next()
			return
		}

		token := tokenParts[1]
		user, err := m.authService.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}

		// Establecer el usuario en el contexto si la autenticación es exitosa
		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("user_role", user.Role)

		c.Next()
	}
}
