package storage

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/Nikby53/image-converter/internal/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Storage holds config and s3 methods.
type Storage struct {
	svc  *s3.S3
	conf *configs.AWSConfig
}

type StorageInterface interface {
	UploadFile(image io.ReadSeeker, fileID string) error
	UploadTargetFile(filename, fileID string) error
	DownloadFile(fileID string) ([]byte, error)
	DownloadImageFromID(fileID string) (string, error)
}

func connectToAws(conf *configs.AWSConfig) (*session.Session, error) {
	s3session, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(conf.AccID, conf.SecretKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("can't create session, %w", err)
	}
	return s3session, nil
}

func initS3ServiceClient(conf *configs.AWSConfig) (*s3.S3, error) {
	s3session, err := connectToAws(conf)
	if err != nil {
		return nil, err
	}

	return s3.New(s3session), nil
}

// New is constructor for Storage.
func New(conf *configs.AWSConfig) (*Storage, error) {
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

// UploadTargetFile uploads target file to aws s3 bucket.
func (s *Storage) UploadTargetFile(filename, fileID string) error {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	_, err = s.svc.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(s.conf.BucketName),
		Key:    aws.String(fileID),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return fmt.Errorf("can't upload file: %w", err)
	}

	return nil
}

// DownloadFile downloads file from aws storage.
func (s *Storage) DownloadFile(fileID string) ([]byte, error) {
	resp, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.conf.BucketName),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return nil, fmt.Errorf("can't download file with id %s: %w", fileID, err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't serialize response body: %w", err)
	}

	return buf, nil
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
