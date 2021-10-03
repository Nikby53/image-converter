package repository

import (
	"fmt"

	"github.com/Nikby53/image-converter/internal/models"
)

// CreateUser method is for inserting data into users table.
func (r *Repository) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password) values ($1, $2) RETURNING id", users)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("cannot create user: %w", err)
	}
	return id, nil
}

// GetUser gets the user.
func (r *Repository) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", users)
	row := r.db.QueryRow(query, email, password)
	if err := row.Scan(&user.Id); err != nil {
		return models.User{}, fmt.Errorf("cannot find the user in database:%w", err)
	}
	return user, nil
}
