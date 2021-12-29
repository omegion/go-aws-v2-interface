package api

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/s3_mock.go -package=mocks github.com/omegion/go-aws-v2-interface S3Interface
// S3Interface is an interface for API.
type S3Interface interface {
	GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error)
	ListObjectVersions(input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error)
	ListObjects(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	DeleteObjects(input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)
}

// API is main struct of S3.
type API struct {
	Config *aws.Config
	Client *s3.Client
}

// NewAPI inits new API.
func NewAPI() (S3Interface, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &API{
		Config: &cfg,
		Client: client,
	}, nil
}

// GetObject gets object from S3.
func (a API) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return a.Client.GetObject(context.TODO(), input)
}

// ListObjectVersions gets object versions from S3.
func (a API) ListObjectVersions(input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	return a.Client.ListObjectVersions(context.TODO(), input)
}

// ListObjects gets object from S3.
func (a API) ListObjects(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return a.Client.ListObjectsV2(context.TODO(), input)
}

// PutObject puts object to S3.
func (a API) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return a.Client.PutObject(context.TODO(), input)
}

// DeleteObject deletes object from S3.
func (a API) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return a.Client.DeleteObject(context.TODO(), input)
}

// DeleteObjects deletes object from S3.
func (a API) DeleteObjects(input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	return a.Client.DeleteObjects(context.TODO(), input)
}
