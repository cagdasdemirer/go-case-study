package main

import (
	"github.com/gin-gonic/gin"
	"go-case-study/controllers"
	"go-case-study/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/user", controllers.UserCreate)
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
