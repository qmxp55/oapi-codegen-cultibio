package models

type AppConfig struct {
	TimescaleDB map[string]string `mapstructure:"timescaleDB"`
}
