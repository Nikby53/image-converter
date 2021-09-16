package service

import "github.com/Nikby53/image-converter/internal/models"

func (s *Service) UploadImage(image models.Images) (string, error) {
	return s.repoImage.UploadImage(image)
}
