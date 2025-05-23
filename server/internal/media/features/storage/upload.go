package storage

import (
	"fmt"
	"os"
)

func (s *service) Upload(file *os.File, path string) (string, error) {
	if s.s3 == nil {
		return "", fmt.Errorf("no provider")
	}

	return s.s3.Upload(file, path)
}
