package auth

import (
	"fmt"
	"media/internal/errors"
	"media/internal/security"
	"media/internal/util"
	"media/module/user"
)

func (s *service) Register(params user.RegisterParams) (*user.LoginResult, error) {
	userWithEmail, err := s.repository.GetByEmail(params.Email)
	if err != nil {
		return nil, fmt.Errorf("get by email: %w", err)
	}
	if userWithEmail != nil {
		return nil, errors.Publicf("Email '%s' is already used", params.Email)
	}

	hash, err := security.HashPassword(params.Password)
	if err != nil {
		return nil, fmt.Errorf("create hash: %w", err)
	}
	emailToken := util.UniqueRandomString()
	createdUser, err := s.repository.Create(user.CreateParams{
		Name:  params.Name,
		Email: params.Email,
		EmailVerify: user.EmailVerify{
			IsVerified: false,
			Token:      emailToken,
		},
		Password: hash,
		Role:     user.RoleUser,
	})
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	err = s.SendVerifyEmail(createdUser)
	if err != nil {
		return nil, fmt.Errorf("send verify email: %v", err)
	}

	token, err := s.createSession(createdUser)
	if err != nil {
		return nil, fmt.Errorf("creating session: %w", err)
	}

	return &user.LoginResult{User: createdUser, Token: token}, nil
}
