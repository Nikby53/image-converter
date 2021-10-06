package service

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

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
func (s *Service) ConvertImage(imageBytes []byte, targetFormat string, ratio int) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
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

	return buf.Bytes(), nil
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
func (s *Service) GetImageID(id string) (string, error) {
	return s.repoImage.GetImageID(id)
}
