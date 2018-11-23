package main

import (
	"net/http"

	"github.com/spencercharest/plex-collections/app"
	"github.com/spencercharest/plex-collections/routes"
)

func main() {
	application := app.NewApplication()

	defer application.LogFile.Close()
	defer application.Database.Close()

	r := routes.Router(application)
	panic(http.ListenAndServe(":4500", r))
}
