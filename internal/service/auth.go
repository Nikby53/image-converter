package service

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/dgrijalva/jwt-go"
)

var (
	errInvalidSigningMethod = errors.New("invalid signing method")
	errTokenClaimsNotType   = errors.New("token claims are not of type *tokenClaims")
	errWrongPassword        = errors.New("wrong password")
	errGenHashPassword      = errors.New("can't generate hash password")
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 4 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

// CreateUser method creates user.
func (s *Service) CreateUser(ctx context.Context, user models.User) (id int, err error) {
	user.Password, err = generatePasswordHash(user.Password)
	if err != nil {
		return id, errGenHashPassword
	}
	return s.repo.CreateUser(ctx, user)
}

// GenerateToken generates jwt token for user.
func (s *Service) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(context.Background(), email)
	if err != nil {
		return "", err
	}
	if !comparePasswordHash(password, user.Password) {
		return "", errWrongPassword
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
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

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func comparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
