package image

import (
	"image"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/effect"
)

func (s *service) EffectEInk(img image.Image) image.Image {
	img = effect.Grayscale(img)
	img = adjust.Brightness(img, 0.5)
	img = adjust.Contrast(img, 0.5)

	return img
}
