package process

import (
	"fmt"
	"media/module/book"
)

func (s *service) ProcessSourceData(isbn book.ISBN, data book.SourceMultiData) (*book.CreateParams, error) {
	metadata, err := s.NormalizeMetadata(isbn, data)
	if err != nil {
		return nil, fmt.Errorf("normalize metadata: %v", err)
	}

	images := data.Images()
	var bookImages *book.Images
	if len(images) > 0 {
		var err error
		bookImages, err = s.ProcessImage(images[0], fmt.Sprintf("books/%s", metadata.ISBN))
		if err != nil {
			return nil, fmt.Errorf("process images: %v", err)
		}
	}

	processedMetadata, err := s.ProcessMetadata(metadata)
	if err != nil {
		return nil, fmt.Errorf("process metadata: %v", err)
	}

	return &book.CreateParams{
		Metadata: *processedMetadata,
		Sources:  data,
		Images:   *bookImages,
	}, nil
}
