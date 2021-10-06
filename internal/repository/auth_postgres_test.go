package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := New(db)
	tests := []struct {
		name    string
		mock    func()
		input   models.User
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("Test", "password").WillReturnRows(rows)
			},
			input: models.User{
				Email:    "Test",
				Password: "password",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("Test", "").WillReturnRows(rows)
			},
			input: models.User{
				Email:    "Test",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := repo.CreateUser(tt.input)
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

func TestRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		email    string
		password string
	}
	repo := New(db)
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    models.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery("SELECT id FROM users").
					WithArgs("petrov@gmail.com", "13125412312Vv").WillReturnRows(rows)
			},
			input: args{
				email:    "petrov@gmail.com",
				password: "13125412312Vv",
			},
			want: models.User{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name: "Not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT id FROM users").
					WithArgs("petrov@gmail.com", "13125412312Vv").WillReturnRows(rows)
			},
			input: args{
				email:    "petrov@gmail.com",
				password: "13125412312Vv",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := repo.GetUser(tt.input.email, tt.input.password)
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
