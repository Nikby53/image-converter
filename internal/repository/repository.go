package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

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
	GetImageByID(id string) (models.Images, error)
}

type RepositoryInterface interface {
	AuthorizationRepository
	ImagesRepository
	Transactional(f func(repo RepositoryInterface) (string, error)) (string, error)
}

// Repository struct provides access to the database.
type Repository struct {
	db sqlx.Ext
}

// New is constructor of the Repository.
func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Transactional(f func(repo RepositoryInterface) (string, error)) (string, error) {
	sqlDB, ok := r.db.(*sqlx.DB)
	if !ok {
		return "", errors.New("couldn't bring to DB")
	}
	tx, err := sqlDB.Beginx()
	if err != nil {
		return "", fmt.Errorf("couldn't start transaction:%s", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			logrus.Infof("error from(commit) tx %v", err.Error())
		}

	}()

	str, err := f(&Repository{db: tx})
	if err != nil {
		return "", err
	}

	return str, nil
}
