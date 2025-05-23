package image

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/imgio"
)

func (s *service) Open(path string) (image.Image, error) {
	img, err := imgio.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %s", err)
	}

	return img, nil
}
