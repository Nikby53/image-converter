package errors

import "errors"

var (
	// ErrEmailEmpty error is for checking email.
	ErrEmailEmpty = errors.New("email should be not empty")
	// ErrPasswordEmpty error is for checking password.
	ErrPasswordEmpty = errors.New("password should be not empty")
)
