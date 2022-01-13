package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/jmoiron/sqlx"
)

// AuthorizationRepository interface contains database methods of the user.
type AuthorizationRepository interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, email string) (models.User, error)
}

// ImagesRepository interface contains database methods of images.
type ImagesRepository interface {
	InsertImage(ctx context.Context, filename, format string) (string, error)
	RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error)
	GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error)
	UpdateRequest(ctx context.Context, status, imageID, targetID string) error
	GetImageByID(ctx context.Context, id string) (models.Images, error)
}

// RepoInterface contains AuthorizationRepository,
// ImagesRepository and Transactional func.
type RepoInterface interface {
	AuthorizationRepository
	ImagesRepository
	Transactional(f func(repo RepoInterface) error) error
}

// Repository struct provides access to the database.
type Repository struct {
	db sqlx.ExtContext
}

// New is constructor of the Repository.
func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Transactional func implement atomicity,it begins transactions,rollback and commit them.
func (r *Repository) Transactional(f func(repo RepoInterface) error) error {
	sqlDB, ok := r.db.(*sqlx.DB)
	if !ok {
		return errors.New("couldn't bring to DB")
	}

	tx, err := sqlDB.Beginx()
	if err != nil {
		return fmt.Errorf("couldn't start transaction:%w", err)
	}

	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				_ = fmt.Errorf("cannot rollback transaction:%w", err)
				return
			}

			return
		}

		err = tx.Commit()
		if err != nil {
			_ = fmt.Errorf("cannot commit transaction: %w", err)
			return
		}
	}()

	err = f(&Repository{db: tx})
	if err != nil {
		return err
	}

	return nil
}
