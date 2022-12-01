package models

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	UserID        uint   `json:"userID"`
	DeviceID      uint   `json:"deviceID"`
	RelationNotes string `json:"relationNotes"`
	User          User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Device        Device `gorm:"foreignKey:DeviceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
