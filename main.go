package main

import (
	"github.com/gin-gonic/gin"
	"go-case-study/pkg/controllers"
	"go-case-study/pkg/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/user", controllers.UserCreate)
	r.POST("/device", controllers.DeviceCreate)
	r.POST("/relation", controllers.RelationCreate)
	r.DELETE("/relation/:id", controllers.RelationDelete)
	r.GET("/sensor_data", controllers.SensorDataFetch)
	r.GET("/device", controllers.DeviceFetch)
	err := r.Run()
	if err != nil {
		return
	}
}
