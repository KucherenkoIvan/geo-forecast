package main

import (
	http_server "geoforecast/api"
	"geoforecast/internal/config"
	"geoforecast/internal/db"
	"log"
)

func init() {
	config.Load()
	log.Println("[INIT] Config loaded")

	db.Connect()
	log.Println("[INIT] Connected to database")

	db.Migrate()
	log.Println("[INIT] Migration finished")
}

func main() {
	log.Println("[APP] Firing up HTTP server")
	http_server.Start()
}
