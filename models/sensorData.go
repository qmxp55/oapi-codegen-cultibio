package models

import "time"

// SensorData defines model for SensorData.
type SensorData struct {
	SensorID  *int       `json:"sensorid,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Value     *float32   `json:"value,omitempty"`
}

// TableName sets the table name for SensorData.
func (SensorData) TableName() string {
	return "sensordata"
}
