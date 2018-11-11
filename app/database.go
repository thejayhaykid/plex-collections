package app

import (
	"log"

	"github.com/jinzhu/gorm"
	// needed for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// initializeDatabase initializes and returns a sqlite3 database
func initializeDatabase() *gorm.DB {
	databasePath := getDatabaseFilePath()

	db, err := gorm.Open("sqlite3", databasePath)

	if err != nil {
		log.Fatalln("Unable to create or connect to database.")
	}

	return db
}
