package usermod

import (
	"media/internal/cache"
	"media/internal/config"
	"media/internal/email"
	"media/internal/jwt"
	"media/internal/mongo"
	"media/module/user"
	"media/module/user/features/auth"
	"media/module/user/features/repository"
)

func NewRepository(db *mongo.Connection) user.Repository {
	return repository.New(db)
}

func NewAuthService(dashboardConfig config.Dashboard, repository user.Repository, cache *cache.Service, jwt *jwt.Service, emails email.Service) user.AuthService {
	return auth.New(dashboardConfig, repository, cache, jwt, emails)
}
