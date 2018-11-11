package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// getPlexCollectionsPath gets the plex collections path and creates if it does not exist
func getPlexCollectionsPath() string {
	home, err := homedir.Dir()

	if err != nil {
		log.Fatalln("Unable to get user home directory.")
	}

	plexCollectionsPath := filepath.Join(home, ".plex-collections")

	if err := createDirIfNotExists(plexCollectionsPath); err != nil {
		log.Fatalln("Unable to create plex collections path.")
	}

	return plexCollectionsPath
}

// createDirIfNotExists will check a given directory and create it if it doesnt exist
func createDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0700); err != nil {
			return err
		}
	}

	return nil
}
