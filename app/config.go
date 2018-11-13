package app

import (
	"math/rand"
	"time"

	"github.com/spencercharest/plex-collections/models"
)

// Config represents the application config
type Config struct {
	PlexCollectionsPath string
	JWTSecret           string
}

func (a *App) getConfig() {
	config := Config{
		PlexCollectionsPath: getPlexCollectionsPath(),
	}

	settings := models.Settings{}

	a.Database.First(&settings)

	if settings.JWTSecret == "" {
		settings.JWTSecret = createSecret(12)
		a.Database.Save(&settings)
	}

	config.JWTSecret = settings.JWTSecret

	a.Config = config
}

func createSecret(n int) string {
	rand.Seed(time.Now().UnixNano())

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
