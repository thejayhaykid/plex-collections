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

// NewApp runs applications startup tasks and returns an App struct
func NewApp() App {
	app := App{}

	app.getDatabase()
	app.getLogger()
	app.getConfig()

	return app
}
