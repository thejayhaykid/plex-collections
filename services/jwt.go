package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spencercharest/plex-collections/models"
)

// Claims is jwt claims
type Claims struct {
	jwt.StandardClaims
	ID     uint
	Email  string
	Role   string
	Active bool
}

// GenerateJWTToken creates a JWT token based on the user model that was passed in.
func GenerateJWTToken(user models.User, settings models.Settings) string {
	signingKey := []byte(settings.JWTSecret)

	claims := Claims{
		ID:     user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Active: user.Active,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(signingKey)

	return signedToken
}
