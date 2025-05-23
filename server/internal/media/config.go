package media

type ProviderS3Config struct {
	Region          string `json:"region" validate:"required"`
	AccessKeyID     string `json:"accessKeyId" validate:"required"`
	SecretAccessKey string `json:"secretAccessKey" validate:"required"`
	BucketName      string `json:"bucketName" validate:"required"`
	Endpoint        string `json:"endpoint,omitempty" validate:"required"`
	PublicEndpoint  string `json:"publicEndpoint" validate:"required"`
}

type ProvidersConfig struct {
	S3 *ProviderS3Config `json:"s3"`
}

type Config struct {
	TmpFolder string          `json:"tmpFolder" validate:"required"`
	Providers ProvidersConfig `json:"providers" validate:"required"`
}
