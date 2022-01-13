package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/sirupsen/logrus"
)

var (
	errEmailEmpty    = errors.New("email should be not empty")
	errPasswordEmpty = errors.New("password should be not empty")
	errInvalidEmail  = errors.New("invalid email")
	errShortPassword = errors.New("password should has at least 6 letters")
)

// Registration struct holds information about user.
type Registration struct {
	models.User
}

// ValidateSignUp validates signUp handler.
func (r *Registration) ValidateSignUp(req *http.Request) error {
	if r.Email == "" {
		return errEmailEmpty
	}

	if r.Password == "" {
		return errPasswordEmpty
	}

	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegexp.MatchString(r.Email) {
		return errInvalidEmail
	}

	if len(r.Password) < 6 {
		return errShortPassword
	}

	return nil
}

type userID struct {
	ID int `json:"id"`
}

func (s *Server) signUp(w http.ResponseWriter, r *http.Request) {
	var input Registration

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s.errorJSON(w, http.StatusBadRequest, err)

		return
	}

	err = input.ValidateSignUp(r)
	if err != nil {
		s.errorJSON(w, http.StatusBadRequest, err)

		return
	}

	id, err := s.service.CreateUser(r.Context(), input.User)
	if err != nil {
		s.errorJSON(w, http.StatusConflict, err)

		return
	}

	err = json.NewEncoder(w).Encode(userID{ID: id})
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		logrus.Printf("signUp: error encoding json: %v", err)

		return
	}
}

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ValidateSignIn validates signIp handler.
func (r *loginInput) ValidateSignIn(req *http.Request) error {
	if r.Email == "" {
		return errEmailEmpty
	}

	if r.Password == "" {
		return errPasswordEmpty
	}

	return nil
}

type tokenJWT struct {
	Token string `json:"token"`
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var input loginInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s.errorJSON(w, http.StatusBadRequest, err)

		return
	}

	err = input.ValidateSignIn(r)
	if err != nil {
		s.errorJSON(w, http.StatusBadRequest, err)

		return
	}

	token, err := s.service.GenerateToken(r.Context(), input.Email, input.Password)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)

		return
	}

	err = json.NewEncoder(w).Encode(tokenJWT{
		Token: token})
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Printf("login: error encoding json: %v", err)

		return
	}
}
