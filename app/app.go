package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

// App is passed down to controllers and contains application, logger, and database
type App struct {
	Config  Config
	Logger  *logrus.Logger
	LogFile *os.File
}

// Config represents the application config
type Config struct {
	PlexCollectionsPath string
}

// NewApp runs applications startup tasks and returns an App struct
func NewApp() App {
	logger, logFile := newLogger()
	plexCollectionsPath := getPlexCollectionsPath()

	config := Config{
		PlexCollectionsPath: plexCollectionsPath,
	}

	return App{
		Logger:  logger,
		LogFile: logFile,
		Config:  config,
	}
}
