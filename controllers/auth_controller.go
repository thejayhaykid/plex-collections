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

// Signup creates a user in the database
func (c AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	payload := models.UserSignUpPayload{}

	if ok, message := Decode(r, &payload); !ok {
		SendAPIError(w, 400, message)
		return
	}

	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     "user",
		Active:   false,
	}

	if result := c.App.Database.Create(&user); result.Error != nil {
		status, message := parseGormError(result)
		SendAPIError(w, status, message)
		return
	}

	response := models.MapUserDAOtoUserDTO(user)

	SendJSON(w, 200, response)
}

// Signin authenticates a user sign in payload
func (c AuthController) Signin(w http.ResponseWriter, r *http.Request) {
	payload := models.UserSignInPayload{}

	if ok, err := Decode(r, &payload); !ok {
		SendAPIError(w, 400, err)
		return
	}

	user := models.User{Email: strings.ToLower(payload.Email)}

	c.App.Database.Where(user).First(&user)

	if user.ID == 0 {
		SendAPIError(w, 401, "Incorrect email.")
		return
	}

	if !user.ValidPassword(payload.Password) {
		SendAPIError(w, 401, "Incorrect password.")
		return
	}

	if !user.Active {
		SendAPIError(w, 403, "Your account must be activated by your server administrator.")
		return
	}

	settings := models.Settings{}

	if result := c.App.Database.First(&settings); result.Error != nil {
		status, message := parseGormError(result)
		SendAPIError(w, status, message)
		return
	}

	token := services.GenerateJWTToken(user, settings)

	response := models.MapUserDAOToUserAuthenticationResponse(user, token)

	SendJSON(w, 200, response)
}
