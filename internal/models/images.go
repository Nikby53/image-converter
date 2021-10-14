package models

import "os"

// Images information about images.
type Images struct {
	ID     int
	File   *os.File
	Format string
	Name   string
}
