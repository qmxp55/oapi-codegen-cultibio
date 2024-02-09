package dal

import (
	"fmt"
	"oapi-codegen-cultibio/models"
)

func TimeScaleConnectionString(cfg models.AppConfig) string {

	// Access values
	timescaleDB := cfg.TimescaleDB
	host := timescaleDB["host"]
	port := timescaleDB["port"]
	user := timescaleDB["user"]
	password := timescaleDB["password"]
	database := timescaleDB["database"]

	// Use the values as needed in your application
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		user, password, host, port, database)

	return connectionString
}
