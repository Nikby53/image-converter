package repository

import (
	"fmt"

	"github.com/Nikby53/image-converter/internal/models"
)

func (r *Repository) UploadImage(image models.Images) (string, error) {
	var imageID string
	query := fmt.Sprintf("INSERT INTO %s (name, format) VALUES ($1, $2) RETURNING id", images)

	err := r.db.QueryRow(query, image.Name, image.Format).Scan(&imageID)
	if err != nil {
		return "", fmt.Errorf("can't insert image: %w", err)
	}

	return imageID, nil
}
