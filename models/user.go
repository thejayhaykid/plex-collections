package models

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thedevsaddam/govalidator"
	"golang.org/x/crypto/bcrypt"
)

// User represents an application user
type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
}

// BeforeCreate is a hook that runs before a user is created
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return err
	}

	scope.DB().Model(u).Update("email", strings.ToLower(u.Email))
	scope.DB().Model(u).Update("password", string(hash))

	return nil
}

// AfterCreate is a hook that runs after a user is created
func (u *User) AfterCreate(scope *gorm.Scope) error {
	// make first user admin and active
	if u.ID == 1 {
		scope.DB().Model(u).Update("role", "admin")
		scope.DB().Model(u).Update("active", true)
	}

	return nil
}

// ValidPassword returns true if the given password matches the user password hash
func (u *User) ValidPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}

	return true
}

// ParseAndValidate parses a user from a POST request and validated that user
func (u *User) ParseAndValidate(r *http.Request) url.Values {
	rules := govalidator.MapData{
		"firstName": []string{"required"},
		"lastName":  []string{"required"},
		"email":     []string{"required", "email"},
		"password":  []string{"required"},
	}

	messages := govalidator.MapData{
		"firstName": []string{"required:First Name is required."},
		"lastName":  []string{"required:Last Name is required."},
		"email":     []string{"required:Email is required.", "email:Email must be a valid email."},
		"password":  []string{"required:Password is required."},
	}

	opts := govalidator.Options{
		Request:  r,
		Data:     u,
		Rules:    rules,
		Messages: messages,
	}

	v := govalidator.New(opts)

	e := v.ValidateJSON()

	return e
}
