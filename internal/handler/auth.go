package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Nikby53/image-converter/internal/models"
)

var (
	// ErrEmailEmpty error is for checking email.
	errEmailEmpty = errors.New("email should be not empty")
	// ErrPasswordEmpty error is for checking password.
	errPasswordEmpty = errors.New("password should be not empty")
)

type Registration struct {
	models.User
}

func (r *Registration) ValidateSignUp(req *http.Request) error {
	if r.Email == "" {
		return errEmailEmpty
	}
	if r.Password == "" {
		return errPasswordEmpty
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
		log.Printf("signUp: can't decode request body: %v", err)
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
		log.Printf("error signUp: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(userID{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("signUp: error encoding json: %v", err)
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

func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input signInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("signIn: can't decode request body: %v", err)
		return
	}
	err = input.ValidateSignIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := s.services.GenerateToken(input.Email, input.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error signIn: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(tokenJWT{
		Token: token})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("signIn: error encoding json: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
