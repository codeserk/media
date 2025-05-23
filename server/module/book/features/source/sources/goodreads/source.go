package goodreads

import "media/module/book"

type source struct {
	conf *book.SourceGoodreadsConfig
}

func New(conf *book.SourceGoodreadsConfig) book.Source {
	return &source{conf}
}
