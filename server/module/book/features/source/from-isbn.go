package source

import (
	"log"
	"media/module/book"
)

func (s *service) FromISBN(isbn string) (book.SourceMultiData, error) {
	var result []*book.SourceData
	var merged book.Metadata
	for _, source := range s.sources {
		res, err := source.FromISBN(isbn)
		if err != nil {
			log.Printf("error: %v", err)
		}

		if res == nil {
			continue
		}

		result = append(result, res)
		merged.Merge(res.Metadata)
		if merged.IsComplete() {
			return result, nil
		}
	}

	return result, nil
}
