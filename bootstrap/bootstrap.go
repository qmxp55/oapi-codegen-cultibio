package bootstrap

import (
	"fmt"
	"oapi-codegen-cultibio/dal"
	"oapi-codegen-cultibio/models"

	"github.com/spf13/viper"
)

type Application struct {
	TimescaleDB *dal.TimescaleDB
	Config      *models.AppConfig
}

func LoadConfig() (*models.AppConfig, error) {
	viper.SetConfigFile("config/config.yaml")

	var config models.AppConfig
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewInitializeBootsrap() Application {
	// app := Application{}
	// Load configuration
	appConfig, err := LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load configuration: %w", err))
	}

	// Initialize TimescaleDB connection
	timescaleDB := &dal.TimescaleDB{}
	err = timescaleDB.Initialize(dal.TimeScaleConnectionString(*appConfig)) // Use the actual configuration property
	if err != nil {
		panic(fmt.Errorf("failed to initialize TimescaleDB: %w", err))
	}

	return Application{
		TimescaleDB: timescaleDB,
		Config:      appConfig,
	}
}
