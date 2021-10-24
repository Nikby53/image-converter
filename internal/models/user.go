package models

// User contains information about user.
type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
