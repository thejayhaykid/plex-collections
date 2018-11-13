package app

import (
	"log"

	"github.com/jinzhu/gorm"
	// needed for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spencercharest/plex-collections/models"
)

// getDatabase initializes and returns a sqlite3 database
func (a *App) getDatabase() {
	databasePath := getDatabaseFilePath()

	db, err := gorm.Open("sqlite3", databasePath)

	if err != nil {
		log.Fatalln("Unable to create or connect to database.")
	}

	db.AutoMigrate(&models.Settings{})

	a.Database = db
}
