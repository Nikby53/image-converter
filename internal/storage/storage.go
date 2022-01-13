package storage

import (
	"io"
)

// StoragesInterface contains all minio and aws s3 methods.
type StoragesInterface interface {
	UploadFile(image io.ReadSeeker, fileID string) error
	DownloadImageFromID(fileID string) (string, error)
}
