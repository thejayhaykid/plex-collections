package app

// App is passed down to controllers and contains application and database details
type App struct {
	Config Config
}

// Config represents the application config
type Config struct {
	PlexCollectionsPath string
}

// NewApp runs applications startup tasks and returns an App struct
func NewApp() App {
	plexCollectionsPath := getPlexCollectionsPath()

	config := Config{
		PlexCollectionsPath: plexCollectionsPath,
	}

	return App{
		Config: config,
	}
}
