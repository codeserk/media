package source

import "media/module/book"

func (s *service) Search(query string) ([]*book.SourceData, error) {
	for _, source := range s.sources {
		res, err := source.Search(query)
		if err == nil && len(res) > 0 {
			return res, nil
		}
	}

	return nil, nil
}
