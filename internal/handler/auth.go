package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nikby53/image-converter/internal/errors"
	"github.com/Nikby53/image-converter/internal/models"
)

type Registration struct {
	models.User
}

func (r *Registration) ValidateSignUp(req *http.Request) error {
	if r.Email == "" {
		return errors.ErrEmailEmpty
	}
	if r.Password == "" {
		return errors.ErrPasswordEmpty
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
		return errors.ErrEmailEmpty
	}
	if r.Password == "" {
		return errors.ErrPasswordEmpty
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

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) converter(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) request(w http.ResponseWriter, r *http.Request) {

}
