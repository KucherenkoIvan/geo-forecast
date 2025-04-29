package main

import (
	http_server "geoforecast/api"
	"geoforecast/internal/config"
	"geoforecast/internal/db"
)

func init() {
	config.Load()

	db.Connect()

	db.Migrate()
}

func main() {
	http_server.Start()
}
