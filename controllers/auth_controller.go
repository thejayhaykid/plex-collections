package controllers

import (
	"net/http"
	"strings"

	"github.com/spencercharest/plex-collections/app"
	"github.com/spencercharest/plex-collections/models"
	"github.com/spencercharest/plex-collections/services"
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
		SendAPIError(w, 400, message)
		return
	}

	// implicity set these values on user creation
	user.Role = "user"
	user.Active = false

	if result := c.App.Database.Create(&user); result.Error != nil {
		err := result.Error.Error()
		status, message := parseGormError(err)
		SendAPIError(w, status, message)
		return
	}

	SendJSON(w, 200, user)
}

// AuthenticateUser authenticates a user sign in payload
func (c AuthController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	userSignIn := models.UserSignInPayload{}

	if err := userSignIn.ParseAndValidate(r); len(err) != 0 {
		message := getFirstValidationError(err)
		SendAPIError(w, 400, message)
		return
	}

	user := models.User{Email: strings.ToLower(userSignIn.Email)}

	c.App.Database.Where(user).First(&user)

	if user.ID == 0 {
		SendAPIError(w, 401, "Incorrect email.")
		return
	}

	if !user.ValidPassword(userSignIn.Password) {
		SendAPIError(w, 401, "Incorrect password.")
		return
	}

	settings := models.Settings{}

	if result := c.App.Database.First(&settings); result.Error != nil {
		err := result.Error.Error()
		status, message := parseGormError(err)
		SendAPIError(w, status, message)
		return
	}

	token := services.GenerateJWTToken(user, settings)

	response := models.UserResponsePayload{
		Token:  token,
		Role:   user.Role,
		Active: user.Active,
	}

	SendJSON(w, 200, response)
}
