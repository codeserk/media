package lacasadellibro

import "media/module/book"

type source struct {
	conf *book.SourceLaCasaDelLibroConfig
}

func New(conf *book.SourceLaCasaDelLibroConfig) book.Source {
	return &source{conf}
}
