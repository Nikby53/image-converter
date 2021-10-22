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
	return s.repoImage.InsertImage(filename, format)
}

// ConvertImage converts JPG to PNG image and vice versa and compress images with
// the compression ratio specified by the user.
func (s *Service) ConvertImage(sourceImage io.ReadSeeker, targetFormat string, ratio int) (io.ReadSeeker, error) {
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
		var quality *jpeg.Options
		if ratio != 0 {
			quality.Quality = ratio
		}
		if err := jpeg.Encode(buf, img, quality); err != nil {
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

func (s *Service) Convert(payload ConvertPayLoad) (string, error) {
	imageID, err := s.repoImage.InsertImage(payload.Filename, payload.SourceFormat)
	if err != nil {
		return "", fmt.Errorf("can't insert image into db: %w", err)
	}
	err = s.storage.UploadFile(payload.File, imageID)
	if err != nil {
		return "", fmt.Errorf("can't upload file: %w", err)
	}
	sourceFile, err := s.storage.DownloadFile(imageID)
	if err != nil {
		return "", fmt.Errorf("can't download image: %w", err)
	}
	convertedImage, err := s.ConvertImage(sourceFile, payload.TargetFormat, payload.Ratio)
	if err != nil {
		return "", fmt.Errorf("can't convert image: %w", err)
	}
	s.logger.Infof("user with id %v successfully convert image with id %v", payload.UsersID, imageID)
	requestID, err := s.repoImage.RequestsHistory(payload.SourceFormat, payload.TargetFormat, imageID, payload.Filename, payload.UsersID, payload.Ratio)
	if err != nil {
		return "", fmt.Errorf("can't make request: %w", err)
	}
	targetImageID, err := s.repoImage.InsertImage(payload.Filename, payload.TargetFormat)
	if err != nil {
		return "", fmt.Errorf("can't insert image into db: %w", err)
	}
	err = s.repoImage.UpdateRequest(processing, imageID, targetImageID)
	if err != nil {
		return "", fmt.Errorf("can't update status: %w", err)
	}
	err = s.storage.UploadFile(convertedImage, targetImageID)
	if err != nil {
		return "", fmt.Errorf("can't upload image: %w", err)
	}
	err = s.repoImage.UpdateRequest(done, imageID, targetImageID)
	if err != nil {
		return "", fmt.Errorf("can't update status: %w", err)
	}
	return requestID, nil
}

// RequestsHistory inserts history of the users request to the database.
func (s *Service) RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	return s.repoImage.RequestsHistory(sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// GetRequestFromID gets request from user id.
func (s *Service) GetRequestFromID(userID int) ([]models.Request, error) {
	return s.repoImage.GetRequestFromID(userID)
}

// UpdateRequest updates status of request.
func (s *Service) UpdateRequest(status, imageID, targetID string) error {
	return s.repoImage.UpdateRequest(status, imageID, targetID)
}

// GetImageID finds id of the image.
func (s *Service) GetImageByID(id string) (models.Images, error) {
	return s.repoImage.GetImageByID(id)
}
