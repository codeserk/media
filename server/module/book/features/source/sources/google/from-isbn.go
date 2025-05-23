package google

import (
	"fmt"
	"media/module/book"
)

func (s *source) FromISBN(isbn string) (*book.SourceData, error) {
	results, err := s.Search(fmt.Sprintf("isbn:%s", isbn))
	if err != nil {
		return nil, fmt.Errorf("search: %v", err)
	}

	if len(results) == 0 {
		return nil, nil
	}

	return results[0], nil
}
