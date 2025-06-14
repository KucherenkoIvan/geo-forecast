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

type PositionLogRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
	DeviceID  string  `json:"device_id"`
	SessionID string  `json:"session_id"`
}

func PositionLog(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n\n####### ADD POSITION LOG#######\n\n")

	decoder := json.NewDecoder(r.Body)
	var body PositionLogRequest

	err := decoder.Decode(&body)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if body.DeviceID == "" || body.SessionID == "" {
		log.Println("Missing device_id or session_id")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Use provided timestamp or current time if not provided
	timestamp := body.Timestamp
	if timestamp == 0 {
		timestamp = time.Now().UnixMilli()
	}

	rec := models.GeoPositionLog{
		TrackId:   body.SessionID, // Use session ID as track ID
		DeviceId:  body.DeviceID,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
		Timestamp: timestamp,
	}

	if err := db.Connection.Create(&rec).Error; err != nil {
		log.Printf("Error inserting data: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Data inserted, id: %d\n", rec.ID)
	fmt.Fprintln(w, "OK")
	log.Printf("\n\n####### END ADD POSITION LOG #######\n\n")
}
