package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
	defer db.Close()
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
		imagesId     string
		filename     string
		userId       int
		ratio        int
	}
	query := "INSERT INTO request"
	defer db.Close()
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
				imagesId:     "1",
				filename:     "image",
				userId:       1,
				ratio:        54,
			},
			want:    "1",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.RequestsHistory(tt.input.sourceFormat, tt.input.targetFormat, tt.input.imagesId, tt.input.filename, tt.input.userId, tt.input.ratio)
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
