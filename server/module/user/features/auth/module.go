package auth

import (
	"media/internal/cache"
	"media/internal/config"
	"media/internal/email"
	"media/internal/jwt"
	"media/module/user"
)

type service struct {
	dashboardConfig config.Dashboard
	repository      user.Repository
	cache           *cache.Service
	jwt             *jwt.Service
	emails          email.Service
}

func New(dashboardConfig config.Dashboard, repository user.Repository, cache *cache.Service, jwt *jwt.Service, emails email.Service) user.AuthService {
	return &service{dashboardConfig, repository, cache, jwt, emails}
}
