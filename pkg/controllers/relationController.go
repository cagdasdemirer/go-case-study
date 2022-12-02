package controllers

import (
	"github.com/gin-gonic/gin"
	"go-case-study/pkg/initializers"
	"go-case-study/pkg/models"
	"net/http"
)

func RelationCreate(c *gin.Context) {
	var relation models.Relation
	if err := c.BindJSON(&relation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong input !"})
		return
	}

	result := initializers.DB.Create(&relation)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Relation not created."})
		return
	}

	c.IndentedJSON(http.StatusCreated, relation)
}

func RelationDelete(c *gin.Context) {
	var relation models.Relation
	id := c.Param("id")

	result := initializers.DB.First(&relation, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Relation not found."})
		return
	}

	result = initializers.DB.Delete(&relation)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Relation not created."})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"message": "Relation deleted.",
		"data":    relation})
}
