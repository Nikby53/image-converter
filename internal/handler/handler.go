package handler

import (
	"net/http"

	"github.com/Nikby53/image-converter/internal/service"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	services *service.Service
}

func NewServer(service *service.Service) *Server {
	s := Server{
		router:   mux.NewRouter(),
		services: service,
	}
	s.router.HandleFunc("/sign-up", s.SignUp).Methods("POST")
	s.router.HandleFunc("/sign-in", s.SignIn).Methods("POST")
	s.router.HandleFunc("/logout", s.logout).Methods("GET")
	s.router.HandleFunc("/request", s.request).Methods("GET")
	s.router.HandleFunc("/converter", s.converter).Methods("POST")
	return &s
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
