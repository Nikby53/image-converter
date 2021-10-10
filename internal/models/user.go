package models

// User contains information about user.
//
// A user is the security principal for this application.
//
// swagger:model User
type User struct {
	// the ID for this user
	//
	// required: false
	ID int `json:"id" db:"id"`
	// the Email for this user
	//
	// required: true
	Email string `json:"email"`
	// the password for this user
	//
	// required: true
	Password string `json:"password"`
}
