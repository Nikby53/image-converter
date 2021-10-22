package storage

import (
	"io"
)

type StorageInterface interface {
	UploadFile(image io.ReadSeeker, fileID string) error
	UploadTargetFile(filename, fileID string) error
	DownloadFile(fileID string) (io.ReadSeeker, error)
	DownloadImageFromID(fileID string) (string, error)
}
