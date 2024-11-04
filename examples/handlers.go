package main

import (
	"log/slog"
	"net/http"

	"github.com/angelofallars/htmx-go"

	"github.com/joshuar/go-templ-daisyui-examples/templates"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request) {
	// Check, if the current URL is '/'.
	if r.URL.Path != "/" {
		// If not, return HTTP 404 error.
		http.NotFound(w, r)
		slog.Error("render page", "method", r.Method, "status", http.StatusNotFound, "path", r.URL.Path)
		return
	}

	// Define template meta tags.
	metaTags := templates.MetaTags(
		"go-templ-daisyui, example site",                        // define meta keywords
		"Generate beautiful UIs with DaisyUI using Templ + Go.", // define meta description
	)

	// Define template layout for index page.
	pageTemplate := templates.Page(
		"Welcome to the go-templ-daisyui Example!", // define title text
		metaTags, // define meta tags
	)

	// Render index page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, pageTemplate); err != nil {
		// Send HTTP 500 error with log.
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}
