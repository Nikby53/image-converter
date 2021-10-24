package repository

import (
	"database/sql"
	"fmt"

	"github.com/Nikby53/image-converter/internal/models"
)

// InsertImage inserts image to the database and returns image id.
func (r *Repository) InsertImage(filename, format string) (string, error) {
	var imageID string
	query := fmt.Sprintf("INSERT INTO %s (name, format) VALUES ($1, $2) RETURNING id", images)

	err := r.db.QueryRowx(query, filename, format).Scan(&imageID)
	if err != nil {
		return "", fmt.Errorf("can't insert image: %w", err)
	}

	return imageID, nil
}

// RequestsHistory add data to request table and returns request id.
func (r *Repository) RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	var requestID string
	query := fmt.Sprintf("INSERT INTO %s (sourceformat, targetFormat,image_id,filename,user_id, ratio,status) VALUES ($1, $2, $3, $4,$5, $6, 'queued') RETURNING id", request)
	err := r.db.QueryRowx(query, sourceFormat, targetFormat, imageID, filename, userID, ratio).Scan(&requestID)
	if err != nil {
		return "", fmt.Errorf("can't insert request: %w", err)
	}

	return requestID, nil
}

// UpdateRequest updates request status.
func (r *Repository) UpdateRequest(status, imageID, targetID string) error {
	query := fmt.Sprintf("UPDATE %s SET status =$1, target_id=$3 WHERE image_id =$2", request)
	_, err := r.db.Exec(query, status, imageID, targetID)
	if err != nil {
		return fmt.Errorf("can't update status: %w", err)
	}
	return nil
}

// GetRequestFromID allows to get the history of users requests.
func (r *Repository) GetRequestFromID(userID int) ([]models.Request, error) {
	var requestModel []models.Request
	query := fmt.Sprintf("SELECT created, updated, sourceformat, targetformat,status, ratio, filename, image_id, target_id FROM %s WHERE user_id=$1;", request)
	rows, _ := r.db.Query(query, userID)
	requests := models.Request{}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
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

// GetImageByID gets id of the image.
func (r *Repository) GetImageByID(id string) (models.Images, error) {
	var image models.Images
	query := fmt.Sprintf("SELECT id, name, format FROM %s WHERE id=$1", images)
	err := r.db.QueryRowx(query, id).Scan(&image.ID, &image.Name, &image.Format)
	if err != nil {
		return models.Images{}, fmt.Errorf("can't get id: %w", err)
	}
	return image, nil
}
