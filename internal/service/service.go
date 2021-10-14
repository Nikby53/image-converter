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

// Images contains methods for images.
type Images interface {
	InsertImage(filename, format string) (string, error)
	ConvertImage(imageBytes []byte, targetFormat string, ratio int) ([]byte, error)
	RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error)
	GetRequestFromID(userID int) ([]models.Request, error)
	UpdateRequest(status, imageID, targetID string) error
	GetImageID(id string) (string, error)
	GetImage(id string) (name, format string)
}

// ServicesInterface holds Authorization and Images interfaces.
type ServicesInterface interface {
	Authorization
	Images
}

// Service contains repository interfaces.
type Service struct {
	repo      repository.AuthorizationRepository
	repoImage repository.ImagesRepository
}

// New is constructor for Service.
func New(repos repository.AuthorizationRepository, reposImages repository.ImagesRepository) *Service {
	return &Service{
		repo:      repos,
		repoImage: reposImages,
	}
}
