package storage

import (
	"media/internal/media"
	"media/internal/media/features/storage/providers/s3"
)

type service struct {
	conf *media.Config
	s3   *s3.Service
}

func New(conf *media.Config) media.StorageService {
	var s3Service *s3.Service
	if conf.Providers.S3 != nil {
		s3Service = s3.New(conf.Providers.S3)
	}

	return &service{conf, s3Service}
}
