package main

import (
	"github.com/spencercharest/plex-collections/app"
)

func main() {
	application := app.NewApp()
	defer application.LogFile.Close()
}
