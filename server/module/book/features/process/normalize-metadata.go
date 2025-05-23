package process

import (
	"fmt"
	"media/internal/util"
	"media/module/book"
)

func (s *service) NormalizeMetadata(isbn book.ISBN, data book.SourceMultiData) (*book.Metadata, error) {
	metadata := data.Metadata()
	metadata.ISBN = isbn
	if !metadata.IsValid() {
		return nil, fmt.Errorf("invalid metadata")
	}

	metadata.Title = util.Capitalize(metadata.Title)
	metadata.Publisher = util.Capitalize(metadata.Publisher)

	return &metadata, nil
}
