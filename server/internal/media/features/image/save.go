package image

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
)

func (s *service) Save(img image.Image, path string) (string, error) {
	if !strings.HasSuffix(path, ".png") {
		path += ".png"
	}

	savePath, err := s.getFilePath()
	if err != nil {
		return "", fmt.Errorf("get file path: %v", err)
	}
	err = imgio.Save(savePath, img, imgio.PNGEncoder())
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}
	savedFile, err := os.Open(savePath)
	if err != nil {
		return "", fmt.Errorf("open saved file: %v", err)
	}
	defer savedFile.Close()

	return s.storage.Upload(savedFile, path)
}
