package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represents an application user at the database level
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
	Active   bool
}

// BeforeCreate is a hook that runs before a user is created
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return err
	}

	u.Email = strings.ToLower(u.Email)
	u.Password = string(hash)

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
