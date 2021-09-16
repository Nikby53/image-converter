package service

import (
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/repository"
)

// Authorization contains methods for authorization of a user.
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Images interface {
	UploadImage(image models.Images) (string, error)
}

// Service contains interfaces.
type Service struct {
	repo      repository.AuthorizationRepository
	repoImage repository.ImagesRepository
}

// New is Service constructor.
func New(repos repository.AuthorizationRepository, reposImages repository.ImagesRepository) *Service {
	return &Service{
		repo:      repos,
		repoImage: reposImages,
	}
}
