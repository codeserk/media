package google

import (
	"context"
	"fmt"
	"media/module/book"

	"regexp"
	"time"

	"github.com/samber/lo"
	"google.golang.org/api/books/v1"
	"google.golang.org/api/option"
)

func (s *source) Search(query string) ([]*book.SourceData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := books.NewService(ctx, option.WithAPIKey(s.conf.ApiKey))
	if err != nil {
		return nil, fmt.Errorf("new service: %v", err)
	}

	booksService := books.NewVolumesService(client)
	result, err := booksService.List(query).LangRestrict("es").Do()
	if err != nil {
		return nil, fmt.Errorf("search books: %v", err)
	}

	return lo.Map(result.Items, func(item *books.Volume, _ int) *book.SourceData {
		return toMetadata(item)
	}), nil
}

var searchFieldRegex = regexp.MustCompile("[^\\d\\w,\\(\\)]*")
var searchFields string = searchFieldRegex.ReplaceAllString(`
	items(
		id,
		saleInfo(buyLink, saleability),
		volumeInfo(authors, categories, imageLinks(thumbnail), industryIdentifiers, infoLink, language, pageCount, printType, publishedDate, title)
	)
`, "")
var searchFilter string = ""
