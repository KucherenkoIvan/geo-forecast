package controllers

import (
	"encoding/json"
	"fmt"
	"geoforecast/internal/db"
	"geoforecast/internal/db/models"
	"log"
	"net/http"
	"time"
)

func PositionLog(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n\n####### ADD POSITION LOG#######\n\n")

	log.Printf("get track id from query\n")
	trackId := r.URL.Query().Get("trackId")

	if trackId == "" {
		log.Println("got empty trackId, not inserting\n")
		fmt.Fprintln(w, "OK")
		return
	}

	log.Printf("end get track id from query, id: %s\n", trackId)

	decoder := json.NewDecoder(r.Body)

	var body struct {
		Latitude  float64
		Longitude float64
	}

	err := decoder.Decode(&body)

	log.Printf("body parse end, error == nil = %t\n", err == nil)

	if err == nil {
		timestamp := time.Now().UnixMilli()

		log.Printf("got device id\n")

		rec := models.GeoPositionLog{
			TrackId:   trackId,
			Latitude:  body.Latitude,
			Longitude: body.Longitude,
			Timestamp: timestamp,
		}

		db.Connection.Create(&rec)

		log.Printf("data inserted, id: %d\n", rec.ID)
	}

	fmt.Fprintln(w, "OK")
	log.Printf("\n\n####### END ADD POSITION LOG #######\n\n")
}
