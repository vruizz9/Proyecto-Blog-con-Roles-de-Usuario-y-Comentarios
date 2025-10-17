package domain

type Role string

const (
	RoleAdmin Role = "Administrador"
	RoleUser  Role = "Usuario"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}
