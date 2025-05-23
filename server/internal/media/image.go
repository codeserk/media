package media

import "image"

type ResizeFilter int

const (
	Nearest ResizeFilter = iota
	Linear
)

type ResizeParams struct {
	Width           int
	Height          int
	KeepAspectRatio bool
	Filter          ResizeFilter
}

type ImageService interface {
	Open(path string) (image.Image, error)
	Download(url string) (string, error)
	Resize(img image.Image, params ResizeParams) (image.Image, error)

	// Effects
	Pixelate(img image.Image) (image.Image, error)
	EffectEInk(img image.Image) image.Image

	Base64(img image.Image) (string, error)
	Save(img image.Image, path string) (string, error)
}
