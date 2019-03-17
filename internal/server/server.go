package server

import (
	"log"
	"net/http"
	"time"

	"github.com/PeterBooker/colligo/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	//"github.com/mholt/certmagic"
)

// Server holds all the data the App needs
type Server struct {
	Logger *log.Logger
	Config *config.Config
	Router *chi.Mux
}

// New returns a pointer to the main server struct
func New(l *log.Logger, c *config.Config) *Server {
	s := &Server{
		Config: c,
		Logger: l,
	}

	return s
}

// Setup starts the HTTP Server
func (s *Server) Setup() {
	//err := certmagic.HTTPS([]string{"colligo.dev"}, s.Router)
	//if err != nil {
	//s.Logger.Fatalf("HTTP server failed to start: %s\n")
	//}

	s.Router = chi.NewRouter()

	// Middleware Stack
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.DefaultCompress)
	s.Router.Use(middleware.RedirectSlashes)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(15 * time.Second))

	s.routes()

	http := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      s.Router,
		Addr:         ":80",
	}
	go func() { s.Logger.Fatal(http.ListenAndServe()) }()
}
