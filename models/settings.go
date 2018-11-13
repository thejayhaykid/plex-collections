package models

import "github.com/jinzhu/gorm"

// Settings represent saved application settings
type Settings struct {
	gorm.Model
	JWTSecret string
}
