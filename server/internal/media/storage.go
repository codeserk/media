package media

import "os"

type StorageService interface {
	Upload(file *os.File, path string) (string, error)
}
