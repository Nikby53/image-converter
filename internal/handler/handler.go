package handler

import (
	"net/http"

	"github.com/Nikby53/image-converter/internal/storage"

	"github.com/Nikby53/image-converter/internal/service"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	services service.ServiceInterface
	storage  *storage.Storage
}

func NewServer(service service.ServiceInterface, storage *storage.Storage) *Server {
	s := Server{
		router:   mux.NewRouter(),
		services: service,
		storage:  storage,
	}
	s.router.HandleFunc("/sign-up", s.signUp).Methods("POST")
	s.router.HandleFunc("/login", s.login).Methods("POST")
	api := s.router.NewRoute().Subrouter()
	api.Use(s.UserIdentity)
	api.HandleFunc("/convert", s.convert).Methods("POST")
	api.HandleFunc("/requestHistory", s.requestHistory).Methods("GET")

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
