package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router returns the application mux
func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))

	return r
}
