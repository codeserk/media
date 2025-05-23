package auth

import (
	"fmt"
	"media/module/user"
)

func (s *service) GetUserFromToken(t string) (*user.Entity, error) {
	claims, err := s.jwt.GetClaimsFromToken(t)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	u, err := s.repository.GetById(claims.Id)
	if err != nil {
		return nil, fmt.Errorf("get by id: %w", err)
	}
	if u == nil {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
