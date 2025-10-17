# Blog Backend - Arquitectura Hexagonal

Este es el backend del proyecto de blog implementado con arquitectura hexagonal (también conocida como arquitectura de puertos y adaptadores) en Go.

## 🏗️ Arquitectura

La aplicación sigue los principios de la arquitectura hexagonal:

```
┌─────────────────────────────────────────────────────────────┐
│                    ADAPTERS (EXTERNAL)                     │
├─────────────────────────────────────────────────────────────┤
│  HTTP API  │  Database  │  JWT Auth  │  Config  │  Logger │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    PORTS (INTERFACES)                      │
├─────────────────────────────────────────────────────────────┤
│ UserRepository │ BlogRepository │ CommentRepository │ AuthService │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                  SERVICES (APPLICATION)                    │
├─────────────────────────────────────────────────────────────┤
│ UserService │ AuthService │ BlogService │ CommentService │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    DOMAIN (CORE)                          │
├─────────────────────────────────────────────────────────────┤
│   User   │   Blog   │  Comment  │  Errors  │   Roles   │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Estructura del Proyecto

```
backend/
├── cmd/                           # Punto de entrada
│   └── server/
│       └── main.go               # Wiring y configuración
├── internal/                      # Lógica interna del dominio
│   ├── domain/                    # Entidades y errores
│   │   ├── user.go               # Entidad Usuario
│   │   ├── blog.go               # Entidad Blog
│   │   ├── comment.go            # Entidad Comentario
│   │   └── errors.go             # Errores de dominio
│   ├── ports/                     # Interfaces (puertos)
│   │   ├── user_repository.go    # UserRepository
│   │   ├── blog_repository.go    # BlogRepository
│   │   ├── comment_repository.go # CommentRepository
│   │   └── auth_service.go       # AuthService
│   └── services/                  # Casos de uso
│       ├── user_service.go       # Gestión de usuarios
│       ├── auth_service.go       # Autenticación
│       ├── blog_service.go       # Gestión de blogs
│       └── comment_service.go    # Gestión de comentarios
├── adapters/                      # Adaptadores externos
│   ├── persistence/               # Implementaciones de repositorios
│   │   ├── user_repo_sql.go      # UserRepository con SQL
│   │   ├── blog_repo_sql.go      # BlogRepository con SQL
│   │   ├── comment_repo_sql.go   # CommentRepository con SQL
│   │   └── migrations/           # Esquemas de BD
│   ├── api/                       # API HTTP
│   │   └── http/
│   │       ├── handlers/          # Controladores HTTP
│   │       ├── middleware/        # Middleware de autenticación
│   │       └── router.go          # Configuración de rutas
│   ├── auth/                      # Autenticación JWT
│   │   └── jwt_service.go         # Implementación de AuthService
│   └── config/                    # Configuración
│       └── config.go              # Carga de configuración
└── pkg/                           # Utilidades
    └── logger.go                  # Logger personalizado
```

## 🚀 Instalación

### Prerrequisitos

- Go 1.21 o superior
- MySQL/MariaDB
- Git

### Pasos de instalación

1. **Clonar el repositorio**
   ```bash
   git clone <repository-url>
   cd backend
   ```

2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Configurar base de datos**
   ```bash
   # Crear base de datos y tablas
   mysql -u root -p < adapters/persistence/migrations/schema.sql
   ```

4. **Configurar variables de entorno**
   ```bash
   # Copiar archivo de ejemplo
   cp .env.example .env
   
   # Editar .env con tus configuraciones
   nano .env
   ```

5. **Ejecutar la aplicación**
   ```bash
   go run cmd/server/main.go
   ```

## ⚙️ Configuración

### Variables de Entorno

| Variable | Descripción | Valor por defecto |
|----------|-------------|-------------------|
| `SERVER_HOST` | Host del servidor | `0.0.0.0` |
| `SERVER_PORT` | Puerto del servidor | `8080` |
| `DB_HOST` | Host de la base de datos | `localhost` |
| `DB_PORT` | Puerto de la base de datos | `3306` |
| `DB_USER` | Usuario de la base de datos | `root` |
| `DB_PASSWORD` | Contraseña de la base de datos | - |
| `DB_NAME` | Nombre de la base de datos | `blog_db` |
| `JWT_SECRET_KEY` | Clave secreta para JWT | `your-secret-key` |

## 🔐 Autenticación

La aplicación utiliza JWT (JSON Web Tokens) para la autenticación.

### Endpoints de Autenticación

- `POST /api/auth/register` - Registro de usuarios
- `POST /api/auth/login` - Inicio de sesión
- `GET /api/auth/profile` - Perfil del usuario (requiere autenticación)
- `PUT /api/auth/change-password` - Cambio de contraseña (requiere autenticación)

### Uso de Tokens

Incluir el token en el header de las peticiones:
```
Authorization: Bearer <your-jwt-token>
```

## 📚 API Endpoints

### Usuarios (Administradores)
- `GET /api/admin/users` - Listar todos los usuarios
- `GET /api/admin/users/:id` - Obtener usuario por ID
- `PUT /api/admin/users/:id` - Actualizar usuario
- `DELETE /api/admin/users/:id` - Eliminar usuario

### Blogs
- `GET /api/blogs` - Listar todos los blogs (público)
- `GET /api/blogs/:id` - Obtener blog por ID (público)
- `GET /api/blogs/author/:authorId` - Blogs por autor (público)
- `POST /api/blogs` - Crear blog (requiere autenticación)
- `PUT /api/blogs/:id` - Actualizar blog (autor o admin)
- `DELETE /api/blogs/:id` - Eliminar blog (autor o admin)

### Comentarios
- `GET /api/blogs/:blogId/comments` - Comentarios de un blog (público)
- `POST /api/blogs/:blogId/comments` - Crear comentario (requiere autenticación)
- `PUT /api/comments/:id` - Actualizar comentario (autor o admin)
- `DELETE /api/comments/:id` - Eliminar comentario (autor o admin)

## 🗄️ Base de Datos

### Tablas

- **users**: Usuarios del sistema
- **blogs**: Entradas del blog
- **comments**: Comentarios en los blogs

### Usuarios por Defecto

- **admin** / admin123 (Administrador)
- **usuario1** / user123 (Usuario)
- **usuario2** / user123 (Usuario)

## 🧪 Testing

```bash
# Ejecutar tests
go test ./...

# Ejecutar tests con coverage
go test -cover ./...
```

## 🚀 Despliegue

### Docker

```bash
# Construir imagen
docker build -t blog-backend .

# Ejecutar contenedor
docker run -p 8080:8080 --env-file .env blog-backend
```

### Producción

1. Cambiar `GIN_MODE` a `release`
2. Usar una clave JWT segura y única
3. Configurar HTTPS
4. Configurar logs apropiados
5. Configurar monitoreo y métricas

## 📝 Logs

La aplicación incluye logging estructurado para:
- Información del servidor
- Errores de base de datos
- Operaciones de autenticación
- Accesos a la API

## 🔧 Desarrollo

### Agregar Nueva Funcionalidad

1. **Definir entidad en `internal/domain/`**
2. **Crear interfaz en `internal/ports/`**
3. **Implementar servicio en `internal/services/`**
4. **Crear repositorio en `adapters/persistence/`**
5. **Implementar handler en `adapters/api/http/handlers/`**
6. **Agregar rutas en `adapters/api/http/router.go`**

### Estructura de Commits

```
feat: agregar funcionalidad de búsqueda de blogs
fix: corregir validación de roles en middleware
docs: actualizar README con nuevos endpoints
refactor: reorganizar servicios de autenticación
```

## 🤝 Contribución

1. Fork el proyecto
2. Crear una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 🆘 Soporte

Si tienes problemas o preguntas:

1. Revisar la documentación
2. Buscar en issues existentes
3. Crear un nuevo issue con detalles del problema
4. Contactar al equipo de desarrollo

## 🔄 Changelog

### v1.0.0
- Implementación inicial con arquitectura hexagonal
- Sistema de autenticación JWT
- CRUD completo para usuarios, blogs y comentarios
- Sistema de roles (Administrador/Usuario)
- API REST con Gin
- Base de datos MySQL/MariaDB
