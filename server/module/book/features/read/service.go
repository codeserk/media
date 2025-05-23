package read

import "media/module/book"

type service struct {
	repository book.Repository
	source     book.SourceService
	process    book.ProcessService
}

func New(repository book.Repository, source book.SourceService, process book.ProcessService) book.ReadService {
	return &service{repository, source, process}
}
