package service

import (
	"github.com/Nikby53/image-converter/internal/storage"

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
	RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error)
	GetRequestFromID(userID int) ([]models.Request, error)
	UpdateRequest(status, imageID, targetID string) error
	GetImageByID(id string) (models.Images, error)
	Convert(payload ConvertPayLoad) (string, error)
}

// ServicesInterface holds Authorization and Images interfaces.
type ServicesInterface interface {
	Authorization
	Images
}

// Service contains repository interfaces.
type Service struct {
	repo    repository.RepositoryInterface
	storage *storage.Storage
}

// New is constructor for Service.
func New(repo repository.RepositoryInterface, storageAWS *storage.Storage) *Service {
	return &Service{
		repo:    repo,
		storage: storageAWS,
	}
}
