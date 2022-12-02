package controllers

import (
	"github.com/gin-gonic/gin"
	"go-case-study/pkg/initializers"
	"go-case-study/pkg/models"
	"net/http"
)

func DeviceCreate(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong input !"})
		return
	}

	result := initializers.DB.Create(&device)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Device not created."})
		return
	}

	c.IndentedJSON(http.StatusCreated, device)
}

func DeviceFetch(c *gin.Context) {
	var device []models.Device

	id, controlID := c.GetQuery("id")
	dType, controlType := c.GetQuery("type")
	name, controlName := c.GetQuery("name")

	result := initializers.DB.Find(&device)

	if controlID {
		result = initializers.DB.First(&device, id)
	} else {
		if controlType {
			if controlName {
				result = initializers.DB.Where("device_type = ? AND device_name = ?", dType, name).Find(&device)
			} else {
				result = initializers.DB.Where("device_type = ?", dType).Find(&device)
			}
		} else if controlName {
			result = initializers.DB.Where("device_name = ?", name).Find(&device)
		}
	}

	if result.Error != nil || result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Device not found."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": device})
}
