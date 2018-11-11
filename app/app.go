package app

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// App is passed down to controllers and contains application, logger, and database
type App struct {
	Config   Config
	Logger   *logrus.Logger
	LogFile  *os.File
	Database *gorm.DB
}

// Config represents the application config
type Config struct {
	PlexCollectionsPath string
}

// NewApp runs applications startup tasks and returns an App struct
func NewApp() App {
	plexCollectionsPath := getPlexCollectionsPath()
	logger, logFile := initializeLogger()
	database := initializeDatabase()

	config := Config{
		PlexCollectionsPath: plexCollectionsPath,
	}

	return App{
		Config:   config,
		Logger:   logger,
		LogFile:  logFile,
		Database: database,
	}
}
