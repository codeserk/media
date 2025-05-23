package jwt

import (
	"media/internal/config"
)

type Service struct {
	conf config.JWT
}

func New(conf config.JWT) *Service {
	return &Service{conf}
}
