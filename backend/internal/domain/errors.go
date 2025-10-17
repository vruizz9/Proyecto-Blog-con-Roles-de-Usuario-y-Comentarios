package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("usuario no encontrado")
	ErrUserAlreadyExists  = errors.New("el usuario ya existe")
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrUnauthorized       = errors.New("no autorizado")
	ErrForbidden          = errors.New("acceso prohibido")
	ErrBlogNotFound       = errors.New("blog no encontrado")
	ErrCommentNotFound    = errors.New("comentario no encontrado")
	ErrInvalidInput       = errors.New("entrada inválida")
)
