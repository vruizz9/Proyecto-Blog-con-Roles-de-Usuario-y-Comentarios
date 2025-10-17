package httprouter

import (
	"blog-backend/adapters/api/http/handlers"
	"blog-backend/adapters/api/http/middleware"
	"blog-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// Router configura todas las rutas de la aplicación
type Router struct {
	userHandler    *handlers.UserHandler
	authHandler    *handlers.AuthHandler
	blogHandler    *handlers.BlogHandler
	commentHandler *handlers.CommentHandler
	authMiddleware *middleware.AuthMiddleware
}

// NewRouter crea una nueva instancia del router
func NewRouter(
	userService *services.UserService,
	authService *services.AuthService,
	blogService *services.BlogService,
	commentService *services.CommentService,
	authMiddleware *middleware.AuthMiddleware,
) *Router {
	return &Router{
		userHandler:    handlers.NewUserHandler(userService),
		authHandler:    handlers.NewAuthHandler(authService),
		blogHandler:    handlers.NewBlogHandler(blogService),
		commentHandler: handlers.NewCommentHandler(commentService),
		authMiddleware: authMiddleware,
	}
}

// SetupRoutes configura todas las rutas de la aplicación
func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Middleware global
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Rutas públicas
	public := router.Group("/api")
	{
		// Autenticación
		public.POST("/auth/login", r.authHandler.Login)
		public.POST("/auth/register", r.userHandler.Register)

		// Blogs públicos (solo lectura)
		public.GET("/blogs", r.blogHandler.ListBlogs)
		public.GET("/blogs/author/:authorId", r.blogHandler.GetBlogsByAuthor)
		public.GET("/blogs/:id", r.blogHandler.GetBlog)

		// Comentarios públicos (solo lectura)
		public.GET("/blogs/:id/comments", r.commentHandler.GetCommentsByBlog)
	}

	// Rutas protegidas (requieren autenticación)
	protected := router.Group("/api")
	protected.Use(r.authMiddleware.Authenticate())
	{
		// Perfil de usuario
		protected.GET("/auth/profile", r.authHandler.GetProfile)
		protected.PUT("/auth/change-password", r.authHandler.ChangePassword)

		// Gestión de blogs (autenticados)
		protected.POST("/blogs", r.blogHandler.CreateBlog)
		protected.PUT("/blogs/:id", r.blogHandler.UpdateBlog)
		protected.DELETE("/blogs/:id", r.blogHandler.DeleteBlog)

		// Gestión de comentarios (autenticados)
		protected.POST("/blogs/:id/comments", r.commentHandler.CreateComment)
		protected.PUT("/comments/:id", r.commentHandler.UpdateComment)
		protected.DELETE("/comments/:id", r.commentHandler.DeleteComment)
	}

	// Rutas de administración (solo administradores)
	admin := router.Group("/api/admin")
	admin.Use(r.authMiddleware.Authenticate())
	admin.Use(r.authMiddleware.RequireAdmin())
	{
		// Gestión de usuarios (solo administradores)
		admin.GET("/users", r.userHandler.ListUsers)
		admin.GET("/users/:id", r.userHandler.GetUser)
		admin.PUT("/users/:id", r.userHandler.UpdateUser)
		admin.DELETE("/users/:id", r.userHandler.DeleteUser)
	}

	// Ruta de salud
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Servidor funcionando correctamente"})
	})

	return router
}
