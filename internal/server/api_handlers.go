package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PeterBooker/colligo/internal/auth"
	"github.com/go-chi/chi"
)

// errResponse holds data about an error
type errResponse struct {
	Code string `json:"code,omitempty"`
	Err  string `json:"error"`
}

// apiCallback handles users returned from Social sites after they have authenticated
func (s *Server) apiCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := chi.URLParam(r, "service")

		switch service {
		case "facebook":
			s.Logger.Println("No Facebook Service Yet")
		case "twitter":
			auth.TwitterGetToken(*s.Config, w, r)
		default:
			var resp errResponse
			resp.Err = "Social Service not recognised"
			writeJSON(w, resp, http.StatusBadRequest)
		}
	}
}

// apiAuth ...
func (s *Server) apiAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := chi.URLParam(r, "service")

		switch service {
		case "facebook":
			s.Logger.Fatalf("No Facebook Service Yet")
		case "twitter":
			url, err := auth.TwitterSendUser(*s.Config, w, r)
			if err != nil {
				s.Logger.Printf("Failed to obtain token: %s\n", err)
				return
			}
			s.Logger.Printf("URL: %s\n", url)
			http.Redirect(w, r, url, http.StatusSeeOther)
			return
		default:
			var resp errResponse
			resp.Err = "Please provide a valid service"
			writeJSON(w, resp, http.StatusBadRequest)
			return
		}
	}
}

// apiExample ...
func (s *Server) apiExample() http.HandlerFunc {
	type getLoadedResponse struct {
		Loaded bool `json:"loaded"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`example`))
	}
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panicf("Failed to encode JSON: %v\n", err)
	}
}
