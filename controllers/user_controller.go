package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/spencercharest/plex-collections/models"
)

// UserController is a wrapper around all user controllers
type UserController struct {
	DB *gorm.DB
}

// UpdatePermissions handles user permission PUT requests
func (c UserController) UpdatePermissions(w http.ResponseWriter, r *http.Request) {
	payload := models.UserPermissionsPatchPayload{}
	user := models.User{}

	if ok, message := Decode(r, &payload); !ok {
		SendAPIError(w, 400, message)
		return
	}

	if result := c.DB.First(&user, payload.ID); result.Error != nil {
		status, message := parseGormError(result)
		SendAPIError(w, status, message)
		return
	}

	user.Role = payload.Role

	if *payload.Active {
		user.Active = true
	} else {
		user.Active = false
	}

	if result := c.DB.Save(&user); result.Error != nil {
		status, message := parseGormError(result)
		SendAPIError(w, status, message)
		return
	}

	response := models.MapUserDAOtoUserDTO(user)

	SendJSON(w, 200, response)
}
