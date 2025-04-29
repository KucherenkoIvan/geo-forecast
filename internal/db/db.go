package db

import (
	"fmt"
	"geoforecast/internal/config"
	"geoforecast/internal/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB = nil

func Connect() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Values.PG_USER,
		config.Values.PG_PASSWORD,
		config.Values.PG_HOST,
		config.Values.PG_PORT,
		config.Values.PG_DB,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Connection = db
}

func Migrate() {
	Connection.AutoMigrate(&models.GeoPositionLog{})
}
