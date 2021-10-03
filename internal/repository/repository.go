package repository

import (
	"database/sql"

	"github.com/Nikby53/image-converter/internal/models"
)

type AuthorizationRepository interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type ImagesRepository interface {
	InsertImage(filename, format string) (string, error)
	RequestsHistory(sourceFormat, targetFormat, imagesId, filename string, userId, ratio int) (string, error)
	GetRequestFromId(userID int) ([]models.Request, error)
	UpdateRequest(status, imageID, targetID string) error
	GetImageID(id string) (string, error)
}

type Repository struct {
	db *sql.DB
}

// New is the Repository constructor.
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
