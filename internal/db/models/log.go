package models

import "gorm.io/gorm"

type GeoPositionLog struct {
	gorm.Model
	TrackId   string
	Latitude  float64
	Longitude float64
	Timestamp int64
}
