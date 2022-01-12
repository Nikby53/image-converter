package service

import (
	"context"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/Nikby53/image-converter/internal/storage"
)

// Authorization contains methods for authorization of a user.
type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

// Images contains methods for images.
type Images interface {
	InsertImage(ctx context.Context, filename, format string) (string, error)
	RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error)
	GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error)
	UpdateRequest(ctx context.Context, status, imageID, targetID string) error
	GetImageByID(ctx context.Context, id string) (models.Images, error)
	Conversion(ctx context.Context, payload ConversionPayLoad) (string, error)
}

// ServicesInterface holds Authorization and Images interfaces.
type ServicesInterface interface {
	Authorization
	Images
}

// Service contains repository and storages interfaces.
type Service struct {
	repo    repository.RepoInterface
	storage storage.StoragesInterface
}

// New is constructor for Service.
func New(repo repository.RepoInterface, storages storage.StoragesInterface) *Service {
	return &Service{
		repo:    repo,
		storage: storages,
	}
}
