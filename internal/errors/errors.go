package errors

import "errors"

var (
	ErrEmailEmpty    = errors.New("email should be not empty")
	ErrPasswordEmpty = errors.New("password should be not empty")
)
