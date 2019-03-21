package server

import (
	"net/http"
	"time"
)

// App ...
type App struct {
	Name    string
	Version string
	URL     string
}

// Page ...
type Page struct {
	Title       string
	URLPath     string
	Description string
	Time        time.Time
}

// webIndex ...
func (s *Server) webIndex() http.HandlerFunc {

	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		URL:     "https://www.colligo.dev",
	}

	page := Page{
		Title:       "Colligo - Social Data The Easy Way",
		Description: "Imports Social data into WordPress and gives you full control of how it is displayed.",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		page.URLPath = r.URL.Path
		page.Time = time.Now()

		meta := struct {
			Page Page
			App  App
		}{
			page,
			app,
		}

		err := tmpls.ExecuteTemplate(w, "indexPage", meta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// webDocs ...
func (s *Server) webDocs() http.HandlerFunc {

	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		URL:     "https://www.colligo.dev",
	}

	page := Page{
		Title:       "Documentation - Colligo",
		Description: "Documentation on installing and using Colligo.",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		page.URLPath = r.URL.Path
		page.Time = time.Now()

		meta := struct {
			Page Page
			App  App
		}{
			page,
			app,
		}

		err := tmpls.ExecuteTemplate(w, "docsPage", meta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
