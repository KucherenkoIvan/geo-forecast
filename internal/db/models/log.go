package models

import "gorm.io/gorm"

type GeoPositionLog struct {
	gorm.Model
	TrackId   string
	DeviceId  string
	Latitude  float64
	Longitude float64
	Timestamp int64
}
