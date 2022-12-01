package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name"`
	LastName string   `json:"lastName"`
	RoleID   uint     `json:"role"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Company  string   `json:"company"`
	UserRole UserRole `gorm:"foreignKey:RoleID"`
}

type UserRole struct {
	gorm.Model
	RoleName        string `json:"roleName"`
	RoleDescription string `json:"roleDescription"`
}

type Device struct {
	gorm.Model
	DeviceName  string `json:"deviceName"`
	DeviceModel string `json:"deviceModel"`
	DeviceType  string `json:"deviceType"`
}

type Relation struct {
	gorm.Model
	UserID        uint   `json:"userID"`
	DeviceID      uint   `json:"deviceID"`
	RelationNotes string `json:"relationNotes"`
	User          User   `gorm:"foreignKey:UserID"`
	Device        Device `gorm:"foreignKey:DeviceID"`
}

type SensorData struct {
	gorm.Model
	DeviceID   uint    `json:"deviceID"`
	DeviceTemp float64 `json:"deviceTemp"`
	SensorX    float64 `json:"sensorX"`
	SensorY    float64 `json:"sensorY"`
	SensorZ    float64 `json:"sensorZ"`
	Device     Device  `gorm:"foreignKey:DeviceID"`
}
