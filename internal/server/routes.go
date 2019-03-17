package server

import (
	"github.com/go-chi/chi"
)

func (s *Server) routes() {
	// Add Routes
	s.Router.Get("/", s.webExample())
	s.Router.Get("/about", s.webExample())

	// Add API v1 routes
	s.Router.Mount("/api/v1", s.apiRoutes())

	// Handle NotFound
	//s.Router.NotFound(s.notFound())
}

func (s *Server) apiRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/twitter", s.apiExample())

	return r
}
