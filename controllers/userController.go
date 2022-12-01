package controllers

import (
	"github.com/gin-gonic/gin"
	"go-case-study/initializers"
	"go-case-study/models"
	"net/http"
)

func UserCreate(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong input !"})
		return
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not created."})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}
