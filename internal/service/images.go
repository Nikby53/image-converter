package service

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	JPG  = "jpg"
	PNG  = "png"
	JPEG = "jpeg"
)

func (s *Service) InsertImage(filename, format string) (string, error) {
	return s.repoImage.InsertImage(filename, format)
}

func (s *Service) Convert(imageBytes []byte, targetFormat string, ratio int) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.New("unable to decode image")
	}
	buf := new(bytes.Buffer)
	switch targetFormat {
	case PNG:
		var enc png.Encoder
		enc.CompressionLevel = png.CompressionLevel(ratio)
		err := enc.Encode(buf, img)
		if err != nil {
			return nil, errors.New("can't convert in jpg")
		}
	case JPG, JPEG:
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: ratio}); err != nil {
			return nil, errors.New("can't convert in png")
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", targetFormat)
	}

	return buf.Bytes(), nil
}
