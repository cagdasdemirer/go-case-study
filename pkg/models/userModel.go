package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name" binding:"required"`
	LastName string   `json:"lastName" binding:"required"`
	RoleID   uint     `json:"role"`
	Email    string   `json:"email"` // gorm:"unique" can be added after seeding
	Password string   `json:"password" binding:"required"`
	Phone    string   `json:"phone"`
	Company  string   `json:"company" binding:"required"`
	UserRole UserRole `json:"-" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
