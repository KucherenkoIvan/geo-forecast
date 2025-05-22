package controllers

import (
	"encoding/json"
	"geoforecast/internal/db"
	"geoforecast/internal/db/models"
	"log"
	"net/http"
)

func TracksList(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n\n####### GET TRACKS LIST #######\n\n")

	var tracks []models.GeoPositionLog
	db.Connection.Distinct("TrackId").Find(&tracks)

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	encoder.Encode(tracks)

	log.Printf("\n\n####### END GET TRACKS LIST #######\n\n")
}
