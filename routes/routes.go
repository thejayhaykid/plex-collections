package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spencercharest/plex-collections/app"
)

// Router returns the application mux
func Router(application app.App) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))

	return r
}
