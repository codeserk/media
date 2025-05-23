package process

import (
	"fmt"
	"image"
	"image/color"
	"media/internal/media"
	"media/internal/util"
	"media/module/book"

	color_extractor "github.com/marekm4/color-extractor"
	"github.com/samber/lo"
)

var mainImageSize = 512
var mediumImageSize = 128
var smallImageSize = 32

func (s *service) ProcessImage(url string, basePath string) (*book.Images, error) {
	path, err := s.images.Download(url)
	if err != nil {
		return nil, fmt.Errorf("download image: %v", err)
	}
	img, err := s.images.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open image: %v", err)
	}

	main, mainUrl, err := s.generateImage(img, mainImageSize, fmt.Sprintf("%s/main", basePath))
	if err != nil {
		return nil, fmt.Errorf("generate main: %v", err)
	}

	medium, mediumUrl, err := s.generateImage(img, mediumImageSize, fmt.Sprintf("%s/medium", basePath))
	if err != nil {
		return nil, fmt.Errorf("generate medium: %v", err)
	}
	_, smallUrl, err := s.generateImage(img, smallImageSize, fmt.Sprintf("%s/small", basePath))
	if err != nil {
		return nil, fmt.Errorf("generate small: %v", err)
	}

	pixelatedBase64, err := s.generatePixelatedImage(main)
	if err != nil {
		return nil, fmt.Errorf("generate pixelated image: %v", err)
	}
	inkUrl, err := s.generateInkImage(medium, basePath)
	if err != nil {
	}

	colors := lo.Map(color_extractor.ExtractColors(main), func(item color.Color, _ int) string {
		return util.ColorToString(item)
	})

	return &book.Images{
		MainUrl:     mainUrl,
		MediumUrl:   mediumUrl,
		SmallUrl:    smallUrl,
		InkUrl:      inkUrl,
		PixelBase64: pixelatedBase64,
		Colors:      colors,
	}, nil
}

func (s *service) generatePixelatedImage(img image.Image) (string, error) {
	pixelated, err := s.images.Pixelate(img)
	if err != nil {
		return "", fmt.Errorf("pixelate: %v", err)
	}
	pixelatedBase64, err := s.images.Base64(pixelated)
	if err != nil {
		return "", fmt.Errorf("base64: %v", err)
	}

	return pixelatedBase64, nil
}

func (s *service) generateImage(img image.Image, size int, path string) (image.Image, string, error) {
	processed, err := s.images.Resize(img, media.ResizeParams{Width: size, Height: 0, KeepAspectRatio: true, Filter: media.Linear})
	if err != nil {
		return nil, "", fmt.Errorf("resize: %v", err)
	}
	url, err := s.images.Save(processed, path)
	if err != nil {
		return nil, "", fmt.Errorf("save: %v", err)
	}

	return processed, url, nil
}

func (s *service) generateInkImage(img image.Image, path string) (string, error) {
	processed := s.images.EffectEInk(img)
	url, err := s.images.Save(processed, fmt.Sprintf("%s/ink", path))
	if err != nil {
		return "", fmt.Errorf("save ink: %v", err)
	}

	return url, nil
}
