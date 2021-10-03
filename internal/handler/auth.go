package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"

	"github.com/Nikby53/image-converter/internal/models"
)

var (
	errEmailEmpty    = errors.New("email should be not empty")
	errPasswordEmpty = errors.New("password should be not empty")
	errInvalidEmail  = errors.New("invalid email")
	errShortPassword = errors.New("password should has at least 6 letters")
)

type Registration struct {
	models.User
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (r *Registration) ValidateSignUp(req *http.Request) error {
	if r.Email == "" {
		return errEmailEmpty
	}
	if r.Password == "" {
		return errPasswordEmpty
	}
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
	w.Header().Set("Content-Type", "application/json")
	var input Registration

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Printf("signUp: can't decode request body: %v", err)
		return
	}
	err = input.ValidateSignUp(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := s.services.CreateUser(input.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Printf("error signUp: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(userID{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Printf("signUp: error encoding json: %v", err)
		return
	}
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *signInInput) ValidateSignIn(req *http.Request) error {
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
	w.Header().Set("Content-Type", "application/json")
	var input signInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Printf("signIn: can't decode request body: %v", err)
		return
	}
	err = input.ValidateSignIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := s.services.GenerateToken(input.Email, input.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("error login: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(tokenJWT{
		Token: token})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Printf("signIn: error encoding json: %v", err)
		return
	}
}
