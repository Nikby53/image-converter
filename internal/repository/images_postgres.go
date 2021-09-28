package repository

import (
	"fmt"

	"github.com/Nikby53/image-converter/internal/models"
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

func (r *Repository) RequestsHistory(sourceFormat, targetFormat, imagesId, filename string, userId, ratio int) (string, error) {
	var requestID string
	query := fmt.Sprintf("INSERT INTO %s (sourceformat, targetFormat,images_id, filename,user_id, ratio) VALUES ($1, $2, $3, $4,$5, $6) RETURNING id", request)
	err := r.db.QueryRow(query, sourceFormat, targetFormat, imagesId, filename, userId, ratio).Scan(&requestID)
	if err != nil {
		return "", fmt.Errorf("can't insert request: %w", err)
	}

	return requestID, nil
}

func (r *Repository) UpdateRequest(status string, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET status =$1 WHERE user_id =$2", request)
	_, err := r.db.Exec(query, status, userId)
	if err != nil {
		return fmt.Errorf("can't update status: %w", err)
	}
	return nil
}

func (r *Repository) GetRequestFromId(userID int) ([]models.Request, error) {
	var requestModel []models.Request
	query := fmt.Sprintf("SELECT created, updated, sourceformat, targetformat,status, ratio, filename FROM %s WHERE user_id=$1;", request)
	rows, _ := r.db.Query(query, userID)
	requests := models.Request{}
	defer rows.Close()
	for rows.Next() {
		r := requests
		err := rows.Scan(&r.Created, &r.Updated, &r.SourceFormat, &r.TargetFormat, &r.Status, &r.Ratio, &r.Filename)
		if err != nil {
			return []models.Request{}, fmt.Errorf("%w", err)
		}
		requestModel = append(requestModel, r)
	}
	return requestModel, nil
}
