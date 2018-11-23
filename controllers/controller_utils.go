package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sort"

	"github.com/jinzhu/gorm"
	"github.com/spencercharest/plex-collections/models"
)

type apiError struct {
	Message string `json:"message"`
}

// SendJSON is an api util to send json to the client
func SendJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(status)
	w.Write(response)
}

// SendAPIError sends an api error message
// note: this will probably be expanded in the future but for now this will just send one user friendly message
func SendAPIError(w http.ResponseWriter, status int, message string) {
	err := apiError{Message: message}
	SendJSON(w, status, err)
}

// Decode unmarshals and validates a JSON payload
func Decode(r *http.Request, model models.RequestPayload) (ok bool, err string) {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return false, "Unable to parse JSON body."
	}

	return model.Validate()
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
func parseGormError(result *gorm.DB) (code int, err string) {
	message := result.Error.Error()

	switch message {

	case "UNIQUE constraint failed: users.email":
		return 400, "An account with this email already exists."

	case "record not found":
		return 400, "Record not found."

	}

	return 500, "An unknown error occurred."
}
