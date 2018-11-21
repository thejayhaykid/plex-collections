package controllers

import (
	"net/http"

	"github.com/spencercharest/plex-collections/app"
	"github.com/spencercharest/plex-collections/models"
)

// UserController is a wrapper around all user controllers
type UserController struct {
	App app.App
}

// UpdatePermissions handles user permission PUT requests
func (c UserController) UpdatePermissions(w http.ResponseWriter, r *http.Request) {
	userPermissionsPatch := models.UserPermissionsPatch{}
	user := models.User{}

	if err := userPermissionsPatch.ParseAndValidate(r); len(err) != 0 {
		message := getFirstValidationError(err)
		SendAPIError(w, 400, message)
		return
	}

	if result := c.App.Database.First(&user, userPermissionsPatch.ID); result.Error != nil {
		err := result.Error.Error()
		status, message := parseGormError(err)
		SendAPIError(w, status, message)
		return
	}

	user.Role = userPermissionsPatch.Role
	user.Active = userPermissionsPatch.Active

	if result := c.App.Database.Save(&user); result.Error != nil {
		err := result.Error.Error()
		status, message := parseGormError(err)
		SendAPIError(w, status, message)
		return
	}

	SendJSON(w, 200, user)
}
