package middleware

import (
	"net/http"

	"github.com/spencercharest/plex-collections/controllers"
)

// AdminRequired is a middleware that ensures request user is an admin
func AdminRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		user := GetUserRequestData(r)

		if user.Role != "admin" {
			controllers.SendAPIError(w, 401, "You don't have permission to perform this action.")
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// ActiveRequired is a middleware that ensures request user is active
func ActiveRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		user := GetUserRequestData(r)

		if !user.Active {
			controllers.SendAPIError(w, 401, "A server administrator must activate your account.")
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
