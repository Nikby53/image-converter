package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
