package main

import (
	"github.com/gin-gonic/gin"
	"go-case-study/pkg/controllers"
	"go-case-study/pkg/initializers"
	"go-case-study/pkg/middlewares"
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
	r.POST("/login", controllers.Login)
	r.DELETE("/relation/:id", middlewares.RequireAuth, controllers.RelationDelete)
	r.GET("/sensor_data", controllers.SensorDataFetch)
	r.GET("/device", controllers.DeviceFetch)
	r.GET("/validate", middlewares.RequireAuth, controllers.UserValidate)
	err := r.Run()
	if err != nil {
		return
	}
}
