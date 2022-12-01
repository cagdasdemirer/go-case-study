package models

import "gorm.io/gorm"

type SensorData struct {
	gorm.Model
	DeviceID   uint    `json:"deviceID"`
	DeviceTemp float64 `json:"deviceTemp"`
	SensorX    float64 `json:"sensorX"`
	SensorY    float64 `json:"sensorY"`
	SensorZ    float64 `json:"sensorZ"`
	Device     Device  `gorm:"foreignKey:DeviceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
