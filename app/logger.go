package app

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// getLogger will instantiate and return a new logger and log file
func (a *App) getLogger() {
	logger := logrus.New()

	logFilePath := getLogFilePath()

	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE, 0700)

	if err != nil {
		log.Fatalln("Unable to create log file.")
	}

	mw := io.MultiWriter(file, os.Stdout)

	logger.SetOutput(mw)

	a.LogFile = file
	a.Logger = logger
}
