package ask

type Credentials struct {
	Username string
	Password string
}

type S3StorageBackend struct {
	S3BucketName       string
	AWSRegionName      string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}
