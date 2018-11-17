package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spencercharest/plex-collections/app"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return k.name
}

var (
	userContextKey = &contextKey{"user"}
)

// UserContext represents user data that will be added to the request context
type UserContext struct {
	ID     uint
	Email  string
	Role   string
	Active bool
}

// JWTMiddleware checks each request for JWT and sets user context on the request
func JWTMiddleware(config app.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := getTokenFromRequest(r)

			if err != nil {
				// TODO: handle error
			}

			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.JWTSecret), nil
			})

			claims, ok := parsedToken.Claims.(jwt.MapClaims)

			if !ok || !parsedToken.Valid {
				// TODO: handle bad token
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, userContextKey, mapJWTClaimsToUserContext(claims))

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

// getTokenFromRequest gets the JWT from a request Authorization header
func getTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Malformed bearer token")
	}

	return authHeaderParts[1], nil
}

// mapJWTClaimsToUserContext maps jwt claims container user data to UserContext
func mapJWTClaimsToUserContext(claims map[string]interface{}) UserContext {

	id := claims["ID"].(float64)

	return UserContext{
		ID:     uint(id),
		Email:  claims["Email"].(string),
		Role:   claims["Role"].(string),
		Active: claims["Active"].(bool),
	}
}

// GetUserRequestData parses user request data from request context
func GetUserRequestData(r *http.Request) UserContext {
	return r.Context().Value(userContextKey).(UserContext)
}
