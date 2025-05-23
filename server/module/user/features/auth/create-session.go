package auth

import (
	"fmt"
	"media/internal/jwt"
	"media/module/user"
	"time"
)

func (s *service) createSession(user *user.Entity) (string, error) {
	return s.createSessionWithDuration(user, sessionDuration)
}

func (s *service) createSessionWithDuration(user *user.Entity, duration time.Duration) (string, error) {
	token, err := s.jwt.GenerateJWT(user.Id, jwt.ClaimTypeUser, sessionDuration)
	if err != nil {
		return "", fmt.Errorf("generating jwt: %w", err)
	}

	if err := s.cache.SetString(fmt.Sprintf("%s/%s", sessionCacheKey, token), token, sessionDuration); err != nil {
		return "", fmt.Errorf("saving jwt in cache: %w", err)
	}

	return token, nil
}
