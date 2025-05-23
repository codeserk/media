package mediamod

import (
	"media/internal/media"
	"media/internal/media/features/image"
	"media/internal/media/features/storage"
)

func NewStorageService(conf *media.Config) media.StorageService {
	return storage.New(conf)
}

func NewImageService(conf *media.Config, storage media.StorageService) media.ImageService {
	return image.New(conf, storage)
}
