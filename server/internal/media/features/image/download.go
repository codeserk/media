package image

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func (s *service) Download(url string) (string, error) {
	res, err := s.getImageFromURL(url)
	if err != nil {
		return "", fmt.Errorf("get image from url: %v", err)
	}
	path, err := s.getFilePath()
	if err != nil {
		return "", fmt.Errorf("get file path: %v", err)
	}
	err = s.saveImageToPath(res, path)
	if err != nil {
		return "", fmt.Errorf("save image to path: %v", err)
	}

	return path, nil
}

func (s *service) getImageFromURL(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get from URL %s", err)
	}

	return res, nil
}

func (s *service) getFilePath() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("generate uuid: %v", err)
	}

	date := time.Now().Format("2006-01-02")
	folder := fmt.Sprintf("%s/%s", s.conf.TmpFolder, date)
	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create tmp folder: %v", err)
	}

	return fmt.Sprintf("%s/%s", folder, id), nil
}

func (r *service) saveImageToPath(res *http.Response, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("copy content to file %v", err)
	}

	return err
}
