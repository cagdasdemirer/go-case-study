package main

import (
	"go-case-study/initializers"
	"go-case-study/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.User{},
		&models.UserRole{},
		&models.Device{},
		&models.Relation{},
		&models.SensorData{})

	if err != nil {
		return
	}
}
