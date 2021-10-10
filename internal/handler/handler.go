package handler

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
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
	// swagger:operation POST /signup signup signup
	// ---
	// summary: Registers a user.
	// description: Could be any user.
	// parameters:
	// - name: user
	//   in: body
	//   description: the user to create
	//   schema:
	//     "$ref": "#/definitions/User"
	// responses:
	//   "201":
	//     description: user created successfully
	//   "401":
	//     description: unauthorized user
	//   "500":
	//     description: internal server error
	s.router.HandleFunc("/login", s.login).Methods("POST")
	// swagger:operation POST /login login login
	// ---
	// summary: Authorizes the user.
	// description: Only authorized user has access.
	// parameters:
	// - name: user
	//   in: body
	//   description: the user to create
	//   schema:
	//     "$ref": "#/definitions/User"
	// responses:
	//   "200":
	//     description: successful operation
	//   "403":
	//     description: not enough right
	//   "500":
	//     description: internal server error
	api := s.router.NewRoute().Subrouter()
	api.Use(s.UserIdentity)
	api.HandleFunc("/convert", s.convert).Methods("POST")
	// swagger:operation POST /convert convert convert
	// ---
	// summary: Create a request to convert an image
	// description: Receives an image from an input form and converts it PNG to JPG and vice versa.
	// security:
	// - bearerAuth: []
	// parameters:
	//  content:
	//	multipart/form-data:
	//  schema:
	//	type: object
	// properties:
	//  file:
	//	type: string
	//  format: binary
	//  source_format:
	//	type: string
	//  enum: [jpeg, png]
	//	target_format:
	//	type: string
	//	enum: [jpeg, png]
	//	ratio:
	//	type: integer
	//	minimum: 1
	//	maximum: 9
	//	description: Image compression ratio
	//   required: true
	// responses:
	//   "200":
	//     description: successful operation
	//   "401":
	//     description: unauthorized user
	//   "500":
	//     description: internal server error
	api.HandleFunc("/requestHistory", s.requestHistory).Methods("GET")
	api.HandleFunc("/downloadImage/{id}", s.downloadImage).Methods("GET")
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
