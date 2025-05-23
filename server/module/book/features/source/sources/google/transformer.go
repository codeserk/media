package google

import (
	"encoding/json"
	"media/module/book"
	"time"

	"google.golang.org/api/books/v1"
)

func toMetadata(c *books.Volume) *book.SourceData {
	originalJSON, _ := json.Marshal(c)

	if c == nil {
		return nil
	}

	publishedAt, err := time.Parse("2006-01-02", c.VolumeInfo.PublishedDate)
	if err != nil {
		publishedAt, _ = time.Parse("2006", c.VolumeInfo.PublishedDate)
	}
	images := []string{}
	if c.VolumeInfo != nil && c.VolumeInfo.ImageLinks != nil {
		images = []string{c.VolumeInfo.ImageLinks.Thumbnail}
	}

	var isbn string
	for _, identifier := range c.VolumeInfo.IndustryIdentifiers {
		if identifier.Type == "ISBN_13" {
			isbn = identifier.Identifier
			break
		}
	}

	return &book.SourceData{
		Source: book.SourceGoogle,
		Metadata: &book.Metadata{
			Title:       c.VolumeInfo.Title,
			Description: c.VolumeInfo.Description,
			Authors:     c.VolumeInfo.Authors,
			ISBN:        book.ToISBN(isbn),
			Publisher:   c.VolumeInfo.Publisher,
			Tags:        c.VolumeInfo.Categories,
			PageCount:   int(c.VolumeInfo.PageCount),
			PublishedAt: publishedAt,
		},
		Images:   images,
		Original: string(originalJSON),
	}
}
