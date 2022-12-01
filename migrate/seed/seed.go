package seed

import (
	"github.com/brianvoe/gofakeit/v6"
	"go-case-study/initializers"
	"go-case-study/models"
	"math/rand"
	"time"
)

func CreateRoles() {
	var roles = []models.UserRole{
		{RoleName: "Admin"},
		{RoleName: "Engineer"},
		{RoleName: "Client"}}

	result := initializers.DB.Create(&roles)
	if result.Error != nil {
		return
	}
}

func CreateUsers() {
	rand.Seed(time.Now().UnixNano())
	var users []models.User
	for i := 0; i < 20; i++ {
		users = append(users, models.User{
			Name:     gofakeit.FirstName(),
			LastName: gofakeit.LastName(),
			RoleID:   uint(rand.Intn(4-1) + 1),
			Email:    gofakeit.Email(),
			Company:  gofakeit.Company()})
	}
	result := initializers.DB.Create(&users)
	if result.Error != nil {
		return
	}
}

func CreateDevices() {
	var devices []models.Device
	for i := 0; i < 50; i++ {
		devices = append(devices, models.Device{
			DeviceName: gofakeit.Animal(),
		})
	}
	result := initializers.DB.Create(&devices)
	if result.Error != nil {
		return
	}
}

func CreateRelations() {
	var relations []models.Relation
	for i := 0; i < 50; i++ {
		relations = append(relations, models.Relation{
			DeviceID: uint(rand.Intn(51-1) + 1),
			UserID:   uint(rand.Intn(21-1) + 1),
		})
	}
	result := initializers.DB.Create(&relations)
	if result.Error != nil {
		return
	}
}

func CreateSensorData() {
	var sensorData []models.SensorData
	for i := 0; i < 200; i++ {
		sensorData = append(sensorData, models.SensorData{
			DeviceID: uint(rand.Intn(51-1) + 1),
			SensorX:  gofakeit.Float64(),
			SensorY:  gofakeit.Float64(),
			SensorZ:  gofakeit.Float64(),
		})
	}
	result := initializers.DB.Create(&sensorData)
	if result.Error != nil {
		return
	}
}

func Seed() {
	CreateRoles()
	CreateUsers()
	CreateDevices()
	CreateRelations()
	CreateSensorData()
}
