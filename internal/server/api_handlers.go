package server

import (
	"net/http"
)

// apiExample ...
func (s *Server) apiExample() http.HandlerFunc {
	type getLoadedResponse struct {
		Loaded bool `json:"loaded"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`example`))
	}
}
