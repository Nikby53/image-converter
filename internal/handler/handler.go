package handler

import (
	"net/http"

	"github.com/Nikby53/image-converter/internal/service"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	services service.ServiceInterface
}

func NewServer(service service.ServiceInterface) *Server {
	s := Server{
		router:   mux.NewRouter(),
		services: service,
	}
	s.router.HandleFunc("/sign-up", s.signUp).Methods("POST")
	s.router.HandleFunc("/login", s.login).Methods("POST")
	api := s.router.NewRoute().Subrouter()
	api.Use(s.UserIdentity)
	api.HandleFunc("/requestHistory", s.requestHistory).Methods("GET")
	api.HandleFunc("/convert", s.convert).Methods("POST")

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
