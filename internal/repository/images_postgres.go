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
	query := fmt.Sprintf("INSERT INTO %s (sourceformat, targetFormat,image_id,filename,user_id, ratio,status) VALUES ($1, $2, $3, $4,$5, $6, 'queued') RETURNING id", request)
	err := r.db.QueryRow(query, sourceFormat, targetFormat, imagesId, filename, userId, ratio).Scan(&requestID)
	if err != nil {
		return "", fmt.Errorf("can't insert request: %w", err)
	}

	return requestID, nil
}

func (r *Repository) UpdateRequest(status, imageID, targetID string) error {
	query := fmt.Sprintf("UPDATE %s SET status =$1, target_id=$3 WHERE image_id =$2", request)
	_, err := r.db.Exec(query, status, imageID, targetID)
	if err != nil {
		return fmt.Errorf("can't update status: %w", err)
	}
	return nil
}

func (r *Repository) GetRequestFromId(userID int) ([]models.Request, error) {
	var requestModel []models.Request
	query := fmt.Sprintf("SELECT created, updated, sourceformat, targetformat,status, ratio, filename, image_id, target_id FROM %s WHERE user_id=$1;", request)
	rows, _ := r.db.Query(query, userID)
	requests := models.Request{}
	defer rows.Close()
	for rows.Next() {
		r := requests
		err := rows.Scan(&r.Created, &r.Updated, &r.SourceFormat, &r.TargetFormat, &r.Status, &r.Ratio, &r.Filename, &r.OriginalImgID, &r.TargetImgID)
		if err != nil {
			return []models.Request{}, fmt.Errorf("%w", err)
		}
		requestModel = append(requestModel, r)
	}
	return requestModel, nil
}

func (r *Repository) GetImageID(id string) (string, error) {
	var imageID string
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", images)
	err := r.db.QueryRow(query, id).Scan(&imageID)
	if err != nil {
		return "", fmt.Errorf("can't get id: %w", err)
	}
	return imageID, nil
}
