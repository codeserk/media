package image

import "media/internal/media"

type service struct {
	conf    *media.Config
	storage media.StorageService
}

func New(conf *media.Config, storage media.StorageService) media.ImageService {
	return &service{conf, storage}
}
