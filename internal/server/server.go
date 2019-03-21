package server

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/PeterBooker/colligo/internal/auth"
	"github.com/PeterBooker/colligo/internal/client"
	"github.com/PeterBooker/colligo/internal/config"
	"github.com/PeterBooker/colligo/internal/templates"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	//"github.com/mholt/certmagic"
)

var (
	tmpls *template.Template
)

// Server holds all the data the App needs
type Server struct {
	Logger *log.Logger
	Config *config.Config
	Router *chi.Mux
	Client *http.Client
}

// New returns a pointer to the main server struct
func New(l *log.Logger, c *config.Config) *Server {

	tmpl := template.New("").Funcs(template.FuncMap{})
	tmpls = template.Must(vfstemplate.ParseGlob(templates.Files, tmpl, "*.html"))

	s := &Server{
		Config: c,
		Logger: l,
		Client: client.New(),
	}

	return s
}

// Setup starts the HTTP Server
func (s *Server) Setup() {
	auth.Setup()

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

	s.startHTTP()
}
