package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "qweqeqsfsdfgderwae"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 4 * time.Hour
)

// AuthService struct is for initialize repository.
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService is constructor of AuthService.
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

// CreateUser method creates user.
func (s *AuthService) CreateUser(user models.User) (id int, err error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken generates jwt token for user.
func (s *AuthService) GenerateToken(email, password string) (string, error) {
	users, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		users.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
