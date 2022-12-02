package models

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	UserID        uint   `json:"userID"`
	DeviceID      uint   `json:"deviceID"`
	RelationNotes string `json:"relationNotes"`
	User          User   `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Device        Device `json:"-" gorm:"foreignKey:DeviceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
