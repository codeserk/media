package s3

import (
	"media/internal/media"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Service struct {
	conf   *media.ProviderS3Config
	client *s3.S3
}

func New(conf *media.ProviderS3Config) *Service {
	conf.PublicEndpoint = strings.TrimSuffix(conf.PublicEndpoint, "/")

	session, _ := session.NewSession()
	client := s3.New(session, &aws.Config{
		Region:      aws.String(conf.Region),
		Endpoint:    &conf.Endpoint,
		Credentials: credentials.NewStaticCredentials(conf.AccessKeyID, conf.SecretAccessKey, ""),
	})

	return &Service{conf, client}
}
