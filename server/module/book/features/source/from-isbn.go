package source

import (
	"log"
	"media/module/book"
)

func (s *service) FromISBN(isbn book.ISBN) (book.SourceMultiData, error) {
	var result []*book.SourceData
	var merged book.Metadata
	for _, source := range s.sources {
		res, err := source.FromISBN(string(isbn))
		if err != nil {
			log.Printf("from isbn (%s): %v", source, err)
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
