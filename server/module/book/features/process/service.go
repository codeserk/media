package process

import (
	"media/internal/media"
	"media/module/ai"
	"media/module/book"
)

type service struct {
	images media.ImageService
	ai     ai.ChatService
}

func New(images media.ImageService, ai ai.ChatService) book.ProcessService {
	return &service{images, ai}
}
