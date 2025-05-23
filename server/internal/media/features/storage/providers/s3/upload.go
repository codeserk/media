package s3

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Service) Upload(file *os.File, path string) (string, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("stat: %v", err)
	}

	input := &s3.PutObjectInput{
		Bucket:        aws.String(s.conf.BucketName),
		Key:           aws.String(path),
		Body:          file,
		ContentLength: aws.Int64(fileInfo.Size()),
		ContentType:   aws.String(http.DetectContentType([]byte{})),
		ACL:           aws.String("public-read"),
	}

	_, err = s.client.PutObject(input)
	if err != nil {
		return "", fmt.Errorf("put object: %v", err)
	}

	return fmt.Sprintf("%s/%s", s.conf.PublicEndpoint, strings.TrimPrefix(path, "/")), nil
}
