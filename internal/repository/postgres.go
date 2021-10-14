package repository

import (
	"database/sql"
	"fmt"

	"github.com/Nikby53/image-converter/internal/configs"
)

const users = "users"
const images = "images"
const request = "request"

// NewPostgresDB gives access for PostgreSQL.
func NewPostgresDB(cfg *configs.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
