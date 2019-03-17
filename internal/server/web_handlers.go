package server

import (
	"net/http"
)

// webExample ...
func (s *Server) webExample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		w.Write([]byte(`example`))
	}
}
