package service

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/Nikby53/image-converter/internal/models"
)

func (s *Service) UploadImage(image models.Images) (string, error) {
	return s.repoImage.UploadImage(image)
}

const (
	JPG  = "jpg"
	PNG  = "png"
	JPEG = "jpeg"
)

func (s *Service) Convert(imageBytes []byte, targetFormat string) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.New("unable to decode image")
	}
	buf := new(bytes.Buffer)
	switch targetFormat {
	case PNG:
		var enc png.Encoder
		err := enc.Encode(buf, img)
		if err != nil {
			return nil, errors.New("can't convert in jpg")
		}
	case JPG, JPEG:
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 1}); err != nil {
			return nil, errors.New("can't convert in png")
		}
	}
	return buf.Bytes(), nil
}
