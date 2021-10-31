package storage

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AWSConfig is config for aws s3 storage.
type AWSConfig struct {
	BucketName string
	AccID      string
	SecretKey  string
	Region     string
}

// Storage holds config and s3 methods.
type Storage struct {
	svc  *s3.S3
	conf *AWSConfig
}

func connectToAws(conf *AWSConfig) (*session.Session, error) {
	s3session, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(conf.AccID, conf.SecretKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("can't create session, %w", err)
	}
	return s3session, nil
}

func initS3ServiceClient(conf *AWSConfig) (*s3.S3, error) {
	s3session, err := connectToAws(conf)
	if err != nil {
		return nil, err
	}

	return s3.New(s3session), nil
}

// New is constructor for Storage.
func New(conf *AWSConfig) (*Storage, error) {
	svc, err := initS3ServiceClient(conf)
	if err != nil {
		return nil, err
	}
	return &Storage{svc: svc, conf: conf}, nil
}

// UploadFile uploads file to aws s3 bucket.
func (s *Storage) UploadFile(image io.ReadSeeker, fileID string) error {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Body:   image,
		Bucket: aws.String(s.conf.BucketName),
		Key:    aws.String(fileID),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return fmt.Errorf("can't upload file: %w", err)
	}

	return nil
}

// DownloadImageFromID downloads image from image id.
func (s *Storage) DownloadImageFromID(fileID string) (string, error) {
	req, _ := s.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.conf.BucketName),
		Key:    aws.String(fileID),
	})
	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("can't create requets's presigned URL, %w", err)
	}

	return url, err
}
