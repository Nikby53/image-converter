package service

import (
	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/repository"
)

// Authorization contains methods for authorization of a user.
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

// Service contains interfaces.
type Service struct {
	Authorization
}

// New is Service constructor.
func New(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
