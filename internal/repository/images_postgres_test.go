package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRepository_InsertImage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		filename string
		format   string
	}
	query := "INSERT INTO images"
	r := New(db)
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    string
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(query).
					WithArgs("image", "jpg").WillReturnRows(rows)
			},
			input: args{
				filename: "image",
				format:   "jpg",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "Empty Fields error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(query).
					WithArgs("image", "").WillReturnRows(rows)
			},
			input: args{
				filename: "image",
				format:   "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.InsertImage(tt.input.filename, tt.input.format)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_RequestsHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		sourceFormat string
		targetFormat string
		imagesID     string
		filename     string
		userID       int
		ratio        int
	}
	query := "INSERT INTO request"
	r := New(db)
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    string
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(query).
					WithArgs("jpg", "png", "1", "image", 1, 54).WillReturnRows(rows)
			},
			input: args{
				sourceFormat: "jpg",
				targetFormat: "png",
				imagesID:     "1",
				filename:     "image",
				userID:       1,
				ratio:        54,
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(query).
					WithArgs("", "png", "1", "image", 1, 54).WillReturnRows(rows)
			},
			input: args{
				sourceFormat: "",
				targetFormat: "png",
				imagesID:     "1",
				filename:     "image",
				userID:       1,
				ratio:        54,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.RequestsHistory(tt.input.sourceFormat, tt.input.targetFormat, tt.input.imagesID, tt.input.filename, tt.input.userID, tt.input.ratio)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_UpdateRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		status   string
		imageID  string
		targetID string
	}
	query := "UPDATE request SET status"
	r := New(db)
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec(query).
					WithArgs("done", "2", "3").WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				status:   "done",
				imageID:  "2",
				targetID: "3",
			},
			wantErr: false,
		},
		{
			name: "Error",
			mock: func() {
				mock.ExpectExec(query).
					WithArgs("2", "3").WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("can't update status")))
			},
			input: args{
				imageID:  "2",
				targetID: "3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := r.UpdateRequest(tt.input.status, tt.input.imageID, tt.input.targetID)
			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestRepository_GetRequestFromID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := "SELECT created, updated, sourceformat, targetformat,status, ratio, filename, image_id, target_id FROM request"
	r := New(db)

	tests := []struct {
		name    string
		mock    func()
		input   int
		want    []models.Request
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"created", "updated", "sourceformat", "targetformat", "status", "ratio", "filename", "image_id", "target_id"}).
					AddRow(time.Time{}, time.Time{}, "png", "jpg", "done", 75, "img", "4", "5")
				mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: []models.Request{{Filename: "img",
				Status:        "done",
				SourceFormat:  "png",
				TargetFormat:  "jpg",
				Ratio:         75,
				Created:       time.Time{},
				Updated:       time.Time{},
				OriginalImgID: "4",
				TargetImgID:   "5"}},
			wantErr: false,
		},
		{
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"created", "updated", "sourceformat", "targetformat", "status", "ratio", "filename", "image_id"}).
					AddRow(time.Time{}, time.Time{}, "png", "jpg", "done", 75, "img", "4")
				mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: []models.Request{{Filename: "img",
				Status:        "done",
				SourceFormat:  "png",
				TargetFormat:  "jpg",
				Ratio:         75,
				Created:       time.Time{},
				Updated:       time.Time{},
				OriginalImgID: "4",
				TargetImgID:   ""}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.GetRequestFromID(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetImageID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := New(db)

	query := "SELECT id FROM images"
	tests := []struct {
		name    string
		mock    func()
		input   string
		want    string
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery(query).
					WithArgs("1").WillReturnRows(rows)
			},
			input:   "1",
			want:    "1",
			wantErr: false,
		},
		{
			name: "Not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(query).
					WithArgs("").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.GetImageID(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetImage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := New(db)
	type args struct {
		name   string
		format string
	}
	query := "SELECT name, format FROM images"
	tests := []struct {
		name    string
		mock    func()
		input   string
		want    args
		wantErr bool
	}{

		{
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "format"})
				mock.ExpectQuery(query).
					WithArgs("").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			_, _, err := r.GetImage(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
