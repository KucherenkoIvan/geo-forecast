package models

import "gorm.io/gorm"

type GeoPositionLog struct {
	gorm.Model
	DeviceId  string
	Latitude  float64
	Longitude float64
	Timestamp int64
}
