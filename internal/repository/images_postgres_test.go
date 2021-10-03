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
	defer db.Close()
	r := New(db)
	tests := []struct {
		name          string
		mock          func()
		filenameInput string
		formatInput   string
		want          string
		wantErr       bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO images").
					WithArgs("image", "jpg").WillReturnRows(rows)
			},
			filenameInput: "image",
			formatInput:   "jpg",
			want:          "1",
		},
		{
			name: "Error",
			mock: func() {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.InsertImage(tt.filenameInput, tt.formatInput)
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
