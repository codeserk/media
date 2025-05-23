package fixture

import (
	"media/internal/jwt"
)

func JWT() *jwt.Service {
	return jwt.New(JWTConfig())
}
