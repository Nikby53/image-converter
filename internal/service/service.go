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
	InsertImage(filename, format string) (string, error)
	ConvertImage(imageBytes []byte, targetFormat string, ratio int) ([]byte, error)
	RequestsHistory(sourceFormat, targetFormat, imagesId, filename string, userId, ratio int) (string, error)
	GetRequestFromId(userID int) ([]models.Request, error)
}

type ServiceInterface interface {
	Authorization
	Images
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
