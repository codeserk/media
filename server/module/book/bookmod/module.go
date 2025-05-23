package bookmod

import (
	"media/internal/media"
	"media/internal/mongo"
	"media/module/ai"
	"media/module/book"
	"media/module/book/features/process"
	"media/module/book/features/read"
	"media/module/book/features/repository"
	"media/module/book/features/source"
)

func NewRepository(db *mongo.Connection) book.Repository {
	return repository.New(db)
}

func NewSourceService(conf *book.SourceConfig) book.SourceService {
	return source.New(conf)
}

func NewProcessService(images media.ImageService, ai ai.ChatService) book.ProcessService {
	return process.New(images, ai)
}

func NewReadService(repository book.Repository, source book.SourceService, process book.ProcessService) book.ReadService {
	return read.New(repository, source, process)
}
