package auth

import (
	"fmt"
	"media/internal/security"
	"media/module/user"
)

func (s *service) Login(params user.LoginParams) (*user.LoginResult, error) {
	u, err := s.repository.GetByEmail(params.Email)
	if err != nil {
		return nil, fmt.Errorf("get by email: %w", err)
	}
	if u == nil {
		return nil, ErrUserNotFound
	}

	if !security.ComparePasswords(u.Password, params.Password) {
		return nil, ErrInvalidCredentials
	}

	session, err := s.createSession(u)
	if err != nil {
		return nil, fmt.Errorf("creating session: %w", err)
	}

	return &user.LoginResult{User: u, Token: session}, nil
}
