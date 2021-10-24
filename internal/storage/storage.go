package storage

import (
	"io"
)

// StoragesInterface contains all aws s3 methods.
type StoragesInterface interface {
	UploadFile(image io.ReadSeeker, fileID string) error
	UploadTargetFile(filename, fileID string) error
	DownloadFile(fileID string) (io.ReadSeeker, error)
	DownloadImageFromID(fileID string) (string, error)
}
