package image

import (
	"image"
	"media/internal/media"

	"github.com/anthonynsimon/bild/transform"
)

var filters = [...]transform.ResampleFilter{transform.NearestNeighbor, transform.Linear}
var resize = transform.Resize

func (s *service) Resize(img image.Image, params media.ResizeParams) (image.Image, error) {
	w, h := s.size(img, params)
	output := resize(img, w, h, resampleFilter(params.Filter))

	return output, nil
}

func (s *service) size(img image.Image, p media.ResizeParams) (int, int) {
	w := p.Width
	h := p.Height
	imgW := img.Bounds().Max.X
	imgH := img.Bounds().Max.Y

	if p.KeepAspectRatio {
		return s.sizeWithAspectRatio(imgW, imgH, w, h)
	}

	if w == 0 {
		w = imgW
	}
	if h == 0 {
		h = imgH
	}

	return w, h
}

func (s *service) sizeWithAspectRatio(imgW int, imgH int, w int, h int) (int, int) {
	aspectRatio := float32(imgW) / float32(imgH)

	if w == 0 {
		if h == 0 {
			return imgW, imgH
		}

		return int(float32(h) * aspectRatio), h
	}

	if h == 0 {
		return w, int(float32(w) / aspectRatio)
	}

	if aspectRatio > 1 {
		adjustedH := int(float32(w) / aspectRatio)
		if adjustedH > h {
			return int(float32(h) * aspectRatio), h
		}

		return w, int(float32(w) / aspectRatio)
	}

	adjustedW := int(float32(h) * aspectRatio)
	if adjustedW > w {
		return w, int(float32(w) / aspectRatio)
	}

	return int(float32(h) * aspectRatio), h
}

func resampleFilter(f media.ResizeFilter) transform.ResampleFilter {
	return filters[f]
}
