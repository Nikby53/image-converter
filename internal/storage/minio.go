package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioConfig is config for minio storage.
type MinioConfig struct {
	BucketName string
	AccID      string
	SecretKey  string
	Region     string
	Endpoint   string
}

// MinioStorage holds config and minio client.
type MinioStorage struct {
	conf        *MinioConfig
	minioClient *minio.Client
}

// NewMinio is constructor for MinioStorage.
func NewMinio(conf *MinioConfig) (*MinioStorage, error) {
	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccID, conf.SecretKey, ""),
		Secure: false,
		Region: conf.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	return &MinioStorage{
		conf:        conf,
		minioClient: minioClient,
	}, nil
}

// UploadFile uploads file to minio bucket.
func (m *MinioStorage) UploadFile(image io.ReadSeeker, fileID string) error {
	_, err := m.minioClient.PutObject(context.Background(), m.conf.BucketName, fileID, image, -1,
		minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to upload file. err: %w", err)
	}

	return nil
}

// DownloadImageFromID downloads image from image id.
func (m *MinioStorage) DownloadImageFromID(fileID string) (string, error) {
	url, err := m.minioClient.PresignedGetObject(context.Background(), m.conf.BucketName, fileID, time.Hour*3, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get file with id: %s, err: %w", fileID, err)
	}

	return url.String(), nil
}
