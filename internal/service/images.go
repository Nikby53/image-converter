package service

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"

	"github.com/Nikby53/image-converter/internal/repository"

	"github.com/Nikby53/image-converter/internal/models"
)

var (
	errUnableToDecode   = errors.New("unable to decode image")
	errCantConvertInJPG = errors.New("can't convert in jpg")
	errCantConvertInPNG = errors.New("can't convert in png")
)

const (
	// JPG is for validation jpg image.
	JPG = "jpg"
	// PNG is for validation png image.
	PNG = "png"
	// JPEG is for validation jpeg image.
	JPEG = "jpeg"
)

// InsertImage inserts image information to database.
func (s *Service) InsertImage(filename, format string) (string, error) {
	return s.repo.InsertImage(filename, format)
}

// ConvertImage converts JPG to PNG image and vice versa and compress images with
// the compression ratio specified by the user.
func (s *Service) ConvertToType(sourceImage io.ReadSeeker, targetFormat string, ratio int) (io.ReadSeeker, error) {
	img, _, err := image.Decode(sourceImage)
	if err != nil {
		return nil, errUnableToDecode
	}
	buf := new(bytes.Buffer)
	switch targetFormat {
	case PNG:
		var enc png.Encoder
		enc.CompressionLevel = png.CompressionLevel(ratio)
		err := enc.Encode(buf, img)
		if err != nil {
			return nil, errCantConvertInJPG
		}
	case JPG, JPEG:
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: ratio}); err != nil {
			return nil, errCantConvertInPNG
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", targetFormat)
	}

	return bytes.NewReader(buf.Bytes()), nil
}

const (
	processing = "processing"
	done       = "done"
)

type ConvertPayLoad struct {
	SourceFormat string
	TargetFormat string
	Filename     string
	Ratio        int
	File         multipart.File
	UsersID      int
}

func (s *Service) Conversion(payload ConvertPayLoad) (string, error) {
	convertedImage, err := s.ConvertToType(payload.File, payload.TargetFormat, payload.Ratio)
	if err != nil {
		return "", fmt.Errorf("can't convert image: %w", err)
	}
	requestID, err := s.repo.Transactional(func(repo repository.RepositoryInterface) (string, error) {
		imageID, err := repo.InsertImage(payload.Filename, payload.SourceFormat)
		if err != nil {
			return "", fmt.Errorf("can't insert image into db: %w", err)
		}
		targetImageID, err := repo.InsertImage(payload.Filename, payload.TargetFormat)
		if err != nil {
			return "", fmt.Errorf("can't insert image into db: %w", err)
		}
		requestID, err := repo.RequestsHistory(payload.SourceFormat, payload.TargetFormat, imageID, payload.Filename, payload.UsersID, payload.Ratio)
		if err != nil {
			return "", fmt.Errorf("can't make request: %w", err)
		}
		err = repo.UpdateRequest(processing, imageID, targetImageID)
		if err != nil {
			return "", fmt.Errorf("can't update status: %w", err)
		}
		err = repo.UpdateRequest(done, imageID, targetImageID)
		if err != nil {
			return "", fmt.Errorf("can't update status: %w", err)
		}
		err = s.storage.UploadFile(payload.File, imageID)
		if err != nil {
			return "", fmt.Errorf("can't upload file: %w", err)
		}
		err = s.storage.UploadFile(convertedImage, targetImageID)
		if err != nil {
			return "", fmt.Errorf("can't upload image: %w", err)
		}
		return requestID, nil
	})
	if err != nil {
		return "", err
	}
	return requestID, nil
}

// RequestsHistory inserts history of the users request to the database.
func (s *Service) RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	return s.repo.RequestsHistory(sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// GetRequestFromID gets request from user id.
func (s *Service) GetRequestFromID(userID int) ([]models.Request, error) {
	return s.repo.GetRequestFromID(userID)
}

// UpdateRequest updates status of request.
func (s *Service) UpdateRequest(status, imageID, targetID string) error {
	return s.repo.UpdateRequest(status, imageID, targetID)
}

// GetImageID finds id of the image.
func (s *Service) GetImageByID(id string) (models.Images, error) {
	return s.repo.GetImageByID(id)
}
