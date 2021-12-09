package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/Nikby53/image-converter/internal/configs"
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func dropDB(t *testing.T, db *sqlx.DB) {
	_, err := db.Exec("TRUNCATE TABLE images RESTART IDENTITY CASCADE;")
	if err != nil {
		t.Errorf("Unable to truncate table %v", err)
	}
	_, err = db.Exec("TRUNCATE TABLE request RESTART IDENTITY CASCADE;")
	if err != nil {
		t.Errorf("Unable to truncate table %v", err)
	}
}

func (r *Repository) listImagesTest(ctx context.Context) ([]models.Images, error) {
	var images []models.Images
	query := "SELECT * FROM images;"
	row, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}
	for row.Next() {
		img := models.Images{}
		err := row.Scan(&img.ID, &img.Name, &img.Format)
		if err != nil {
			return nil, fmt.Errorf("error %w", err)
		}
		images = append(images, img)
	}
	return images, nil
}

func TestRepository_Transactional(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	conf := configs.DBConfig{
		Host:     "localhost",
		Port:     "5438",
		Username: "postgres",
		Password: "password",
		DBName:   "imageconverter",
		SSLMode:  "disable",
	}
	db, err := NewPostgresDB(&conf)
	assert.NoError(t, err)
	dropDB(t, db)
	r := New(db)
	t.Run("error", func(t *testing.T) {
		testFunc := func(repo RepoInterface) error {
			_, err := repo.InsertImage(context.Background(), "image.png", "png")
			if err != nil {
				return err
			}
			_, err = repo.InsertImage(context.Background(), "image.jpg", "mp3")
			if err != nil {
				return err
			}
			return nil
		}
		err = r.Transactional(testFunc)
		if err == nil {
			t.Errorf("expected error,got nil")
		}
		images, err := r.listImagesTest(context.Background())
		assert.NoError(t, err)
		expectedRows := 0
		if len(images) != expectedRows {
			t.Errorf("expected %v got: %v", expectedRows, len(images))
		}
	})
	t.Run("no error expected", func(t *testing.T) {
		testFunc := func(repo RepoInterface) error {
			_, err := repo.InsertImage(context.Background(), "image.png", "png")
			if err != nil {
				return err
			}
			_, err = repo.InsertImage(context.Background(), "image.jpg", "jpg")
			if err != nil {
				return err
			}
			return nil
		}
		err = r.Transactional(testFunc)
		if err != nil {
			assert.NoError(t, err)
		}
		images, err := r.listImagesTest(context.Background())
		assert.NoError(t, err)
		expectedRows := 2
		if len(images) != expectedRows {
			t.Errorf("expected %v got: %v", expectedRows, len(images))
		}
	})
}
