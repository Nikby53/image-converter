package models

// User contains information about user.
//
// A user is the security principal for this application.
//
// swagger:model User
type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
