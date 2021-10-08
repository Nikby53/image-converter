package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/gorilla/mux"
)

// Server are complex of routers and services.
type Server struct {
	router     *mux.Router
	services   service.ServicesInterface
	storage    storage.StorageInterface
	httpServer *http.Server
}

// NewServer configures server.
func NewServer(service service.ServicesInterface, storage storage.StorageInterface) *Server {
	s := Server{
		router:   mux.NewRouter(),
		services: service,
		storage:  storage,
	}
	s.router.HandleFunc("/signup", s.signUp).Methods("POST")
	s.router.HandleFunc("/login", s.login).Methods("POST")
	api := s.router.NewRoute().Subrouter()
	api.Use(s.UserIdentity)
	api.HandleFunc("/convert", s.convert).Methods("POST")
	api.HandleFunc("/requestHistory", s.requestHistory).Methods("GET")
	api.HandleFunc("/images/{id}", s.downloadImage).Methods("GET")
	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Run runs the server.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown stops the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}
