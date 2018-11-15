package controllers

import (
	"net/http"

	"github.com/spencercharest/plex-collections/app"
	"github.com/spencercharest/plex-collections/models"
)

// AuthController is a wrapper around auth controllers
type AuthController struct {
	App app.App
}

// CreateUser creates a user in the database
func (c AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	if err := user.ParseAndValidate(r); len(err) != 0 {
		message := getFirstValidationError(err)
		sendAPIError(w, 400, message)
		return
	}

	// implicity set these values on user creation
	user.Role = "user"
	user.Active = false

	if result := c.App.Database.Save(&user); result.Error != nil {
		err := result.Error.Error()
		status, message := parseGormError(err)
		sendAPIError(w, status, message)
		return
	}

	sendJSON(w, 200, user)
}

// AuthenticateUser authenticates a user sign in payload
func (c AuthController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	userSignIn := models.UserSignInPayload{}

	if err := userSignIn.ParseAndValidate(r); len(err) != 0 {
		message := getFirstValidationError(err)
		sendAPIError(w, 400, message)
		return
	}

	user := models.User{Email: userSignIn.Email}

	if result := c.App.Database.Where(user).First(&user); result.Error != nil {
		err := result.Error.Error()
		status, message := parseGormError(err)
		sendAPIError(w, status, message)
		return
	}

	if user.ID == 0 {
		sendAPIError(w, 403, "Incorrect email.")
		return
	}

	if !user.ValidPassword(userSignIn.Password) {
		sendAPIError(w, 403, "Incorrect password.")
		return
	}

	// TODO: send token with response
	sendJSON(w, 200, user)
}
