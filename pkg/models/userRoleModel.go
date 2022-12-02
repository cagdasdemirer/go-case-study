package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleName        string `json:"roleName" binding:"required"`
	RoleDescription string `json:"roleDescription"`
}
