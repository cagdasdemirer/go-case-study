package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name" binding:"required"`
	LastName string   `json:"lastName" binding:"required"`
	RoleID   uint     `json:"role"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Company  string   `json:"company" binding:"required"`
	UserRole UserRole `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
