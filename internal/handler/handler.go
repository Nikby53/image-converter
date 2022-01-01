package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Nikby53/image-converter/internal/logs"
	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// ErrorResponse struct is for json error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Server are complex of routers and services.
type Server struct {
	router     *mux.Router
	services   service.ServicesInterface
	storage    storage.StoragesInterface
	httpServer *http.Server
	logger     *logs.Logger
}

// NewServer configures server.
func NewServer(src service.ServicesInterface, st storage.StoragesInterface) *Server {
	s := Server{
		router:   mux.NewRouter(),
		services: src,
		storage:  st,
		logger:   logs.NewLogger(),
	}
	s.initRouters()
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
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) initRouters() {
	s.router.Use(s.logging)
	s.router.HandleFunc("/auth/signup", s.signUp).Methods("POST")
	s.router.HandleFunc("/auth/login", s.login).Methods("POST")
	api := s.router.NewRoute().Subrouter()
	api.Use(s.userIdentity)
	api.HandleFunc("/image/convert", s.convert).Methods("POST")
	api.HandleFunc("/requests", s.requests).Methods("GET")
	api.HandleFunc("/image/download/{id}", s.downloadImage).Methods("GET")
	s.router.Handle("/swagger.yaml", http.FileServer(http.Dir("api/openapi-spec/")))
	opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	s.router.Handle("/docs", sh)
}

func (s *Server) errorJSON(w http.ResponseWriter, statusCode int, errMsg error) {
	w.WriteHeader(statusCode)
	errRes := ErrorResponse{Error: errMsg.Error()}
	err := json.NewEncoder(w).Encode(&errRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
