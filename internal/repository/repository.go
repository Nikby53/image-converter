package repository

import (
	"database/sql"

	"github.com/Nikby53/image-converter/internal/models"
)

// AuthorizationRepository interface contains database methods of the user.
type AuthorizationRepository interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

// ImagesRepository interface contains database methods of images.
type ImagesRepository interface {
	InsertImage(filename, format string) (string, error)
	RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error)
	GetRequestFromID(userID int) ([]models.Request, error)
	UpdateRequest(status, imageID, targetID string) error
	GetImageID(id string) (string, error)
	GetImage(id string) (name, format string, err error)
}

// Repository struct provides access to the database.
type Repository struct {
	db *sql.DB
}

// New is constructor of the Repository.
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
