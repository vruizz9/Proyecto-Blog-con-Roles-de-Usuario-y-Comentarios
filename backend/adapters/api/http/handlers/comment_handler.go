package handlers

import (
	"blog-backend/internal/domain"
	"blog-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommentHandler maneja las peticiones HTTP relacionadas con comentarios
type CommentHandler struct {
	commentService *services.CommentService
}

// NewCommentHandler crea una nueva instancia del handler de comentarios
func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// CreateCommentRequest define la estructura de la petición de creación de comentario
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// UpdateCommentRequest define la estructura de la petición de actualización de comentario
type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// CreateComment crea un nuevo comentario
func (h *CommentHandler) CreateComment(c *gin.Context) {
	blogIDStr := c.Param("blogId")
	blogID, err := strconv.ParseInt(blogIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de blog inválido"})
		return
	}

	var req CreateCommentRequest
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

	comment, err := h.commentService.CreateComment(blogID, uid, req.Content)
	if err != nil {
		switch err {
		case domain.ErrBlogNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog no encontrado"})
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Comentario creado exitosamente",
		"comment":  comment,
	})
}

// GetCommentsByBlog obtiene todos los comentarios de un blog
func (h *CommentHandler) GetCommentsByBlog(c *gin.Context) {
	blogIDStr := c.Param("blogId")
	blogID, err := strconv.ParseInt(blogIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de blog inválido"})
		return
	}

	comments, err := h.commentService.GetCommentsByBlog(blogID)
	if err != nil {
		switch err {
		case domain.ErrBlogNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog no encontrado"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, comments)
}

// UpdateComment actualiza un comentario existente
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateCommentRequest
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

	comment, err := h.commentService.UpdateComment(id, req.Content, uid, role)
	if err != nil {
		switch err {
		case domain.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		case domain.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para editar este comentario"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Comentario actualizado exitosamente",
		"comment":  comment,
	})
}

// DeleteComment elimina un comentario
func (h *CommentHandler) DeleteComment(c *gin.Context) {
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

	if err := h.commentService.DeleteComment(id, uid, role); err != nil {
		switch err {
		case domain.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		case domain.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para eliminar este comentario"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comentario eliminado exitosamente"})
}
