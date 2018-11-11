package main

import (
	"net/http"

	"github.com/spencercharest/plex-collections/app"
	"github.com/spencercharest/plex-collections/routes"
)

func main() {
	application := app.NewApp()
	defer application.LogFile.Close()

	r := routes.Router()
	panic(http.ListenAndServe(":4500", r))
}
