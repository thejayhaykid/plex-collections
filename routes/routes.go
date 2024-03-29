package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spencercharest/plex-collections/app"
	"github.com/spencercharest/plex-collections/controllers"
)

// Router returns the application mux
func Router(application app.Application) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	authController := controllers.AuthController{DB: application.Database}
	userController := controllers.UserController{DB: application.Database}

	r.Route("/api", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", authController.Signup)
			r.Post("/signin", authController.Signin)
		})

		r.Route("/users", func(r chi.Router) {
			r.Post("/permissions", userController.UpdatePermissions)
		})

	})

	return r
}
