package models

import "regexp"

// UserDTO is a data transer object for an application user
type UserDTO struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
}

// UserSignUpPayload represents a user sign up POST payload
type UserSignUpPayload struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// Validate validates a user sign up POST payload
func (u *UserSignUpPayload) Validate() (ok bool, err string) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if u.Name == "" {
		return false, "Name is required."
	}

	if u.Email == "" {
		return false, "Email is required."
	}

	if !emailRegex.MatchString(u.Email) {
		return false, "Email must be a valid email."
	}

	if u.Password == "" {
		return false, "Password is required."
	}

	if u.ConfirmPassword == "" {
		return false, "Confirm Password is required."
	}

	if u.Password != u.ConfirmPassword {
		return false, "Password must match Confirm Password."
	}

	return true, ""
}

// UserSignInPayload represents a user sign in POST payload
type UserSignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates a user sign in POST payload
func (u *UserSignInPayload) Validate() (ok bool, err string) {
	if u.Email == "" {
		return false, "Email is required."
	}

	if u.Password == "" {
		return false, "Password is required."
	}

	return true, ""
}

// UserPermissionsPatchPayload represents a user permissions PATCH payload
type UserPermissionsPatchPayload struct {
	ID     uint   `json:"id"`
	Role   string `json:"role"`
	Active *bool  `json:"active"`
}

// Validate validates a user permissions PATCH payload
func (u *UserPermissionsPatchPayload) Validate() (ok bool, err string) {
	if u.ID == 0 {
		return false, "ID is required."
	}

	if u.Role == "" {
		return false, "Role is required."
	}

	if u.Active == nil {
		return false, "Active is required."
	}

	return true, ""
}

// UserAuthenticationResponse represents a response after successful user auth
type UserAuthenticationResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
	Token  string `json:"token"`
}
