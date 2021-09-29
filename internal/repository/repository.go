package repository

import (
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type ImagesRepository interface {
	InsertImage(filename, format string) (string, error)
	RequestsHistory(sourceFormat, targetFormat, imagesId, filename string, userId, ratio int) (string, error)
	GetRequestFromId(userID int) ([]models.Request, error)
	UpdateRequest(status string, imageID string) error
}

type Repository struct {
	db *sqlx.DB
}

// New is the Repository constructor.
func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}
