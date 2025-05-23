package google

import (
	"media/module/book"
)

type source struct {
	conf *book.SourceGoogleConfig
}

func New(conf *book.SourceGoogleConfig) book.Source {
	return &source{conf}
}
