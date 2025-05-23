package user

import "net/http"

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}

type LoginResult struct {
	User  *Entity
	Token string
}

type AuthService interface {
	Register(params RegisterParams) (*LoginResult, error)
	Login(params LoginParams) (*LoginResult, error)
	SendVerifyEmail(user *Entity) error
	VerifyEmail(email string, token string) (*Entity, error)

	Middleware(next http.Handler) http.Handler
}
