package source

import (
	"media/module/book"
	"media/module/book/features/source/sources/goodreads"
	"media/module/book/features/source/sources/google"
	lacasadellibro "media/module/book/features/source/sources/la-casa-del-libro"
)

type service struct {
	conf    *book.SourceConfig
	sources []book.Source
}

func New(conf *book.SourceConfig) book.SourceService {
	var sources []book.Source

	if conf.Google.IsEnabled {
		sources = append(sources, google.New(&conf.Google))
	}
	if conf.LaCasaDelLibro.IsEnabled {
		sources = append(sources, lacasadellibro.New(&conf.LaCasaDelLibro))
	}
	if conf.LibraryThing.IsEnabled {
		// sources = append(sources, librarything.New(&conf.LibraryThing))
	}
	if conf.StoryGraph.IsEnabled {
		// sources = append(sources, storygraph.New(&conf.StoryGraph))
	}
	if conf.Goodreads.IsEnabled {
		sources = append(sources, goodreads.New(&conf.Goodreads))
	}

	return &service{conf, sources}
}
