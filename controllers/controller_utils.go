package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
)

type apiError struct {
	Message string `json:"message"`
}

// sendJSON is an api util to send json to the client
func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(status)
	w.Write(response)
}

// sendAPIError sends an api error message
// note: this will probably be expanded in the future but for now this will just send one user friendly message
func sendAPIError(w http.ResponseWriter, status int, message string) {
	err := apiError{Message: message}
	sendJSON(w, status, err)
}

// getFirstValidationError will return the first validation error
func getFirstValidationError(errors url.Values) string {
	keys := make([]string, len(errors))

	i := 0

	for key := range errors {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	firstKey := keys[0]

	return errors[firstKey][0]
}

// parseGormError will parse known gorm errors into user friendly messages
func parseGormError(message string) (code int, err string) {
	switch message {
	case "UNIQUE constraint failed: users.email":
		return 400, "An account with this email already exists."
	}

	return 500, "An unknown error occurred."
}
