package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Nikby53/image-converter/internal/rabbitMQ"

	"github.com/Nikby53/image-converter/internal/logs"

	"github.com/go-openapi/runtime/middleware"

	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/gorilla/mux"
)

// Server are complex of routers and services.
type Server struct {
	router        *mux.Router
	services      service.ServicesInterface
	storage       storage.StorageInterface
	httpServer    *http.Server
	logger        *logs.StandardLogger
	messageBroker *rabbitMQ.Client
}

// NewServer configures server.
func NewServer(service service.ServicesInterface, storage storage.StorageInterface, broker *rabbitMQ.Client) *Server {
	s := Server{
		router:        mux.NewRouter(),
		services:      service,
		storage:       storage,
		logger:        logs.NewLogger(),
		messageBroker: broker,
	}
	s.router.HandleFunc("/user/signup", s.signUp).Methods("POST")
	s.router.HandleFunc("/user/login", s.login).Methods("POST")
	api := s.router.NewRoute().Subrouter()
	api.Use(s.userIdentity)
	api.HandleFunc("/image/convert", s.convert).Methods("POST")
	api.HandleFunc("/requests", s.requests).Methods("GET")
	api.HandleFunc("/image/download/{id}", s.downloadImage).Methods("GET")
	s.router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	s.router.Handle("/docs", sh)

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
