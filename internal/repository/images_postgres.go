package repository

import (
	"fmt"
)

func (r *Repository) InsertImage(filename, format string) (string, error) {
	var imageID string
	query := fmt.Sprintf("INSERT INTO %s (name, format) VALUES ($1, $2) RETURNING id", images)

	err := r.db.QueryRow(query, filename, format).Scan(&imageID)
	if err != nil {
		return "", fmt.Errorf("can't insert image: %w", err)
	}

	return imageID, nil
}
