package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO users"

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := New(sqlxDB)
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
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(query).
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

			got, err := repo.CreateUser(context.Background(), tt.input)
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

	query := "SELECT id, password FROM users"

	type args struct {
		email string
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := New(sqlxDB)
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
				rows := sqlmock.NewRows([]string{"id", "password"}).
					AddRow(1, "12312312321")
				mock.ExpectQuery(query).
					WithArgs("petrov@gmail.com").WillReturnRows(rows)
			},
			input: args{
				email: "petrov@gmail.com",
			},
			want: models.User{
				ID:       1,
				Password: "12312312321",
			},
			wantErr: false,
		},
		{
			name: "Not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(query).
					WithArgs("petrov@gmail.com").WillReturnRows(rows)
			},
			input: args{
				email: "petrov@gmail.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := repo.GetUser(context.Background(), tt.input.email)
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
