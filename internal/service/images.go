package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/repository"
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
func (s *Service) InsertImage(ctx context.Context, filename, format string) (string, error) {
	return s.repo.InsertImage(ctx, filename, format)
}

// ConvertToType converts JPG to PNG image and vice versa and compress images with
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

// ConversionPayLoad is payload for Conversion.
type ConversionPayLoad struct {
	SourceFormat string
	TargetFormat string
	Filename     string
	Ratio        int
	File         multipart.File
	UsersID      int
}

// Conversion func is for all conversion logic.
func (s *Service) Conversion(ctx context.Context, payload ConversionPayLoad) (string, error) {
	convertedImage, err := s.ConvertToType(payload.File, payload.TargetFormat, payload.Ratio)
	if err != nil {
		return "", fmt.Errorf("can't convert image: %w", err)
	}
	requestID, err := s.repo.Transactional(func(repo repository.RepoInterface) (string, error) {
		imageID, err := repo.InsertImage(ctx, payload.Filename, payload.SourceFormat)
		if err != nil {
			return "", fmt.Errorf("can't insert image into db: %w", err)
		}
		targetImageID, err := repo.InsertImage(ctx, payload.Filename, payload.TargetFormat)
		if err != nil {
			return "", fmt.Errorf("can't insert image into db: %w", err)
		}
		requestID, err := repo.RequestsHistory(ctx, payload.SourceFormat, payload.TargetFormat, imageID, payload.Filename, payload.UsersID, payload.Ratio)
		if err != nil {
			return "", fmt.Errorf("can't make request: %w", err)
		}
		err = repo.UpdateRequest(ctx, processing, imageID, targetImageID)
		if err != nil {
			return "", fmt.Errorf("can't update status: %w", err)
		}
		err = repo.UpdateRequest(ctx, done, imageID, targetImageID)
		if err != nil {
			return "", fmt.Errorf("can't update status: %w", err)
		}
		_, err = payload.File.Seek(0, io.SeekStart)
		if err != nil {
			return "", fmt.Errorf("can't reset counter: %w", err)
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
func (s *Service) RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	return s.repo.RequestsHistory(ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// GetRequestFromID gets request from user id.
func (s *Service) GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error) {
	return s.repo.GetRequestFromID(ctx, userID)
}

// UpdateRequest updates status of request.
func (s *Service) UpdateRequest(ctx context.Context, status, imageID, targetID string) error {
	return s.repo.UpdateRequest(ctx, status, imageID, targetID)
}

// GetImageByID get information of image by id.
func (s *Service) GetImageByID(ctx context.Context, id string) (models.Images, error) {
	return s.repo.GetImageByID(ctx, id)
}
