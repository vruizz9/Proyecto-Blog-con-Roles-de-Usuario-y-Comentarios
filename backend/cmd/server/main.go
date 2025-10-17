package main

import (
	httprouter "blog-backend/adapters/api/http"
	"blog-backend/adapters/api/http/middleware"
	"blog-backend/adapters/auth"
	"blog-backend/adapters/config"
	"blog-backend/adapters/persistence"
	"blog-backend/internal/services"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	// Cargar configuración
	cfg := config.Load()

	// Conectar a la base de datos
	db, err := connectDB(cfg.Database)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer db.Close()

	// Verificar conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatalf("Error verificando conexión a la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos establecida exitosamente")

	// Crear repositorios (adaptadores de infraestructura)
	userRepo := persistence.NewUserRepositorySQL(db)
	blogRepo := persistence.NewBlogRepositorySQL(db)
	commentRepo := persistence.NewCommentRepositorySQL(db)

	// Crear servicios de infraestructura
	jwtService := auth.NewJWTService(cfg.JWT.SecretKey)

	// Crear servicios de aplicación (casos de uso)
	userService := services.NewUserService(userRepo, jwtService)
	authService := services.NewAuthService(userRepo, jwtService)
	blogService := services.NewBlogService(blogRepo, userRepo)
	commentService := services.NewCommentService(commentRepo, blogRepo, userRepo)

	// Crear middleware de autenticación
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	// Configurar las rutas usando el router
	router := httprouter.NewRouter(userService, authService, blogService, commentService, authMiddleware)
	ginEngine := router.SetupRoutes()

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	log.Printf("Servidor iniciando en %s", serverAddr)
	log.Printf("API disponible en http://%s/api", serverAddr)
	log.Printf("Endpoint de salud en http://%s/health", serverAddr)

	// Iniciar servidor usando Gin directamente
	if err := ginEngine.Run(serverAddr); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}

// connectDB establece la conexión a la base de datos
func connectDB(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión a la base de datos: %w", err)
	}

	// Configurar pool de conexiones
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return db, nil
}
