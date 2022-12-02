package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	DeviceName  string `json:"deviceName" binding:"required"`
	DeviceModel string `json:"deviceModel"`
	DeviceType  string `json:"deviceType"`
}
