package auth

import (
	"fmt"
	"media/module/user"
)

func (s *service) VerifyEmail(email string, token string) (*user.Entity, error) {
	u, err := s.repository.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("get by email: %v", err)
	}
	if u == nil {
		return nil, fmt.Errorf("user not found")
	}

	if u.EmailVerify.IsVerified {
		return u, nil
	}
	if u.EmailVerify.Token != token {
		return nil, fmt.Errorf("invalid token: %v vs %v", u.Email, token)
	}

	return s.repository.VerifyEmail(u.Id)
}
