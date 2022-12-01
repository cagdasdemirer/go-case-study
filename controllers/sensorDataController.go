package controllers

import (
	"github.com/gin-gonic/gin"
	"go-case-study/initializers"
	"go-case-study/models"
	"net/http"
)

func SensorDataFetch(c *gin.Context) {
	var sensorData []models.SensorData
	result := initializers.DB.Find(&sensorData)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fetch not completed."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":  "Fetch completed.",
		"totalRow": result.RowsAffected,
		"data":     sensorData})
}
