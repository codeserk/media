package book

import (
	"time"
)

type Images struct {
	MainUrl     string
	MediumUrl   string
	SmallUrl    string
	InkUrl      string
	PixelBase64 string
	Colors      []string
}

type Entity struct {
	Id       string
	Metadata Metadata
	Sources  SourceMultiData
	Images   Images

	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Ticker
}
