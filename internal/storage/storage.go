package storage

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Config struct {
	BucketName string
	AccId      string
	SecretKey  string
	Region     string
}

type Storage struct {
	svc  *s3.S3
	conf Config
}

func connectToAws(conf Config) (*session.Session, error) {
	s3session, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(conf.AccId, conf.SecretKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("can't create session, %w", err)
	}
	return s3session, nil
}
func initS3ServiceClient(conf Config) (*s3.S3, error) {
	s3session, err := connectToAws(conf)
	if err != nil {
		return nil, err
	}

	return s3.New(s3session), nil
}
func New(conf Config) (*Storage, error) {
	svc, err := initS3ServiceClient(conf)
	if err != nil {
		return nil, err
	}
	return &Storage{svc: svc, conf: conf}, nil
}
func (s *Storage) UploadFile(file io.ReadSeeker, fileID string) error {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(s.conf.BucketName),
		Key:    aws.String(fileID),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return fmt.Errorf("can't upload file: %w", err)
	}

	return nil
}

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

func (s *Storage) DownloadImageFromID(fileID string) (string, error) {
	req, _ := s.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.conf.BucketName),
		Key:    aws.String(fileID),
	})
	fmt.Println(req)

	return "", nil
}
