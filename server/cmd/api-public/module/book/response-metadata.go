package book

import (
	"media/module/book"
	"time"
)

type MetadataResponse struct {
	Title       string           `json:"title" description:"The title of the book" validate:"required"`
	Description string           `json:"description" description:"A brief description of the book" validate:"required"`
	Authors     []string         `json:"authors" description:"A list of authors of the book" validate:"required"`
	ISBN        book.ISBN        `json:"isbn" description:"The ISBN of the book" validate:"required"`
	EAN         string           `json:"ean" description:"The EAN of the book"`
	Publisher   string           `json:"publisher" description:"The publisher of the book" validate:"required"`
	Genres      book.Genres      `json:"genres" description:"The genres of the book"`
	Themes      book.Themes      `json:"themes" description:"The themes of the book"`
	Moods       book.Moods       `json:"moods" description:"The moods of the book"`
	Settings    book.Settings    `json:"settings" description:"The settings of the book"`
	AgeGroups   book.AgeGroups   `json:"ageGroups" description:"The age groups of the book"`
	PacingTypes book.PacingTypes `json:"pacingTypes" description:"The pacing types of the book"`
	PageCount   int              `json:"pageCount" description:"The number of pages in the book" validate:"required"`
	PublishedAt time.Time        `json:"publishedAt" description:"The publication date of the book" validate:"required"`
}

func ToMetadataResponse(metadata *book.Metadata) MetadataResponse {
	return MetadataResponse{
		Title:       metadata.Title,
		Description: metadata.Description,
		Authors:     metadata.Authors,
		ISBN:        metadata.ISBN,
		EAN:         metadata.EAN,
		Publisher:   metadata.Publisher,
		Genres:      metadata.Genres,
		Themes:      metadata.Themes,
		Moods:       metadata.Moods,
		Settings:    metadata.Settings,
		AgeGroups:   metadata.AgeGroups,
		PacingTypes: metadata.PacingTypes,
		PageCount:   metadata.PageCount,
		PublishedAt: metadata.PublishedAt,
	}
}
