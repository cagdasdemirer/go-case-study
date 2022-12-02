package main

import (
	initializers2 "go-case-study/pkg/initializers"
	models2 "go-case-study/pkg/models"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDB()
}

func main() {
	err := initializers2.DB.AutoMigrate(
		&models2.User{},
		&models2.UserRole{},
		&models2.Device{},
		&models2.Relation{},
		&models2.SensorData{})

	if err != nil {
		return
	}

	// Populate the database with some data

	// seed.Seed()

}
