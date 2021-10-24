package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/dgrijalva/jwt-go"
)

var (
	errInvalidSigningMethod = errors.New("invalid signing method")
	errTokenClaimsNotType   = errors.New("token claims are not of type *tokenClaims")
)

const (
	salt       = "qweqeqsfsdfgderwae"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 4 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

// CreateUser method creates user.
func (s *Service) CreateUser(user models.User) (id int, err error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken generates jwt token for user.
func (s *Service) GenerateToken(email, password string) (string, error) {
	users, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		users.ID,
	})

	return token.SignedString([]byte(signingKey))
}

// ParseToken parses token.
func (s *Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidSigningMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errTokenClaimsNotType
	}

	return claims.ID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
