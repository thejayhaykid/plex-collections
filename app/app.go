package app

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Application is passed down to controllers and contains application, logger, and database
type Application struct {
	Config   Config
	Logger   *logrus.Logger
	LogFile  *os.File
	Database *gorm.DB
}

// NewApplication runs applications startup tasks and returns an App struct
func NewApplication() Application {
	application := Application{}

	application.getDatabase()
	application.getLogger()
	application.getConfig()

	return application
}
