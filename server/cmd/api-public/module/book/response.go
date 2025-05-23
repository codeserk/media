package book

import (
	"media/module/book"
	"time"
)

type Response struct {
	Id         string           `json:"id" description:"Unique identifier of the response" validate:"required"`
	Metadata   MetadataResponse `json:"metadata" description:"Metadata associated with the response" validate:"required"`
	Images     ImagesResponse   `json:"images" description:"Image URLs and colors" validate:"required"`
	CreatedAt  time.Time        `json:"createdAt" description:"Timestamp when the response was created" validate:"required"`
	UpdatedAt  time.Time        `json:"updatedAt" description:"Timestamp when the response was last updated" validate:"required"`
	ArchivedAt *time.Time       `json:"archivedAt,omitempty" description:"Timestamp when the response was archived, if applicable"`
}

func ToResponse(b *book.Entity) Response {
	return Response{
		Id:         b.Id,
		Metadata:   ToMetadataResponse(&b.Metadata),
		Images:     ToImagesResponse(&b.Images),
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
		ArchivedAt: b.ArchivedAt,
	}
}

type ImagesResponse struct {
	MainUrl     string   `json:"mainUrl" description:"URL of the main image" validate:"required"`
	MediumUrl   string   `json:"mediumUrl" description:"URL of the medium-sized image" validate:"required"`
	SmallUrl    string   `json:"smallUrl" description:"URL of the small-sized image" validate:"required"`
	InkUrl      string   `json:"inkUrl" description:"URL of the ink image" validate:"required"`
	PixelBase64 string   `json:"pixelBase64" description:"Base64 encoded pixel data" validate:"required"`
	Colors      []string `json:"colors" description:"List of colors associated with the image" validate:"required"`
}

func ToImagesResponse(i *book.Images) ImagesResponse {
	return ImagesResponse{
		MainUrl:     i.MainUrl,
		MediumUrl:   i.MediumUrl,
		SmallUrl:    i.SmallUrl,
		InkUrl:      i.InkUrl,
		PixelBase64: i.PixelBase64,
		Colors:      i.Colors,
	}
}
