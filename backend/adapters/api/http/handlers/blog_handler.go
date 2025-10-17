package handlers

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BlogHandler maneja las peticiones HTTP relacionadas con blogs
type BlogHandler struct {
	blogService *services.BlogService
}

// NewBlogHandler crea una nueva instancia del handler de blogs
func NewBlogHandler(blogService *services.BlogService) *BlogHandler {
	return &BlogHandler{
		blogService: blogService,
	}
}

// CreateBlogRequest define la estructura de la petición de creación de blog
type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateBlogRequest define la estructura de la petición de actualización de blog
type UpdateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateBlog crea un nuevo blog
func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var req CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Obtener ID del usuario autenticado
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	blog, err := h.blogService.CreateBlog(req.Title, req.Content, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando blog"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog creado exitosamente",
		"blog":    blog,
	})
}

// GetBlog obtiene un blog por su ID
func (h *BlogHandler) GetBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	blog, err := h.blogService.GetBlogByID(id)
	if err != nil {
		switch err {
		case domain.ErrBlogNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog no encontrado"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}

// GetBlogsByAuthor obtiene todos los blogs de un autor
func (h *BlogHandler) GetBlogsByAuthor(c *gin.Context) {
	authorIDStr := c.Param("authorId")
	authorID, err := strconv.ParseInt(authorIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de autor inválido"})
		return
	}

	blogs, err := h.blogService.GetBlogsByAuthor(authorID)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Autor no encontrado"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// ListBlogs lista todos los blogs
func (h *BlogHandler) ListBlogs(c *gin.Context) {
	blogs, err := h.blogService.ListBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// UpdateBlog actualiza un blog existente
func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Obtener información del usuario autenticado
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	userRole, exists := c.Get("user_role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	role, ok := userRole.(domain.Role)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	blog, err := h.blogService.UpdateBlog(id, req.Title, req.Content, uid, role)
	if err != nil {
		switch err {
		case domain.ErrBlogNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog no encontrado"})
		case domain.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para editar este blog"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog actualizado exitosamente",
		"blog":    blog,
	})
}

// DeleteBlog elimina un blog
func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obtener información del usuario autenticado
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	userRole, exists := c.Get("user_role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	uid, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	role, ok := userRole.(domain.Role)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	if err := h.blogService.DeleteBlog(id, uid, role); err != nil {
		switch err {
		case domain.ErrBlogNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog no encontrado"})
		case domain.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para eliminar este blog"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog eliminado exitosamente"})
}
