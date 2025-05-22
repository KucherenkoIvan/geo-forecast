package controllers

import (
	"encoding/json"
	"fmt"
	"geoforecast/internal/db"
	"geoforecast/internal/db/models"
	"log"
	"net/http"
)

func Track(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n\n####### GET TRACK #######\n\n")

	trackId := r.URL.Query().Get("trackId")

	if trackId == "" {
		fmt.Fprintln(w, "invalid trackId")
		w.WriteHeader(400)
		return
	}

	var logs []models.GeoPositionLog
	db.Connection.Find(&logs).Where(&models.GeoPositionLog{TrackId: trackId})

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	encoder.Encode(logs)

	log.Printf("\n\n####### END GET TRACK #######\n\n")
}
