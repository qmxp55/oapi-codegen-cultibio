package dal

import (
	"log"

	"oapi-codegen-cultibio/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// TimescaleDB is a wrapper for GORM TimescaleDB connection.
type TimescaleDB struct {
	DB *gorm.DB
}

// Initialize initializes the TimescaleDB connection.
func (db *TimescaleDB) Initialize(connectionString string) error {
	var err error
	db.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	// Auto-migrate the schema
	db.DB.AutoMigrate(&models.SensorData{})

	return nil
}

// StoreSensorData stores the received sensor data into the TimescaleDB.
func (db *TimescaleDB) AddSensorData(sensorData models.SensorData) error {
	err := db.DB.Create(&sensorData).Error
	if err != nil {
		log.Printf("Error storing sensor data: %v\n", err)
		return err
	}

	return nil
}
