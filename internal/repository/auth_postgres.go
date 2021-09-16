package repository

import (
	"fmt"

	"github.com/Nikby53/image-converter/internal/models"
)

// CreateUser method is for inserting data into users table.
func (a *Repository) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password) values ($1, $2) RETURNING id", users)
	row := a.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetUser gets the user.
func (a *Repository) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", users)
	err := a.db.Get(&user, query, email, password)
	return user, err
}
