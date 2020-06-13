package controllers

import (
	"net/http"

	"farm_tales/models"

	"github.com/gin-gonic/gin"
)

// GET /farms
// Get all farms
func GetFarms(c *gin.Context) {
	var farms []models.Farm
	models.DB.Preload("FarmAddress").Find(&farms)
	c.JSON(http.StatusOK, gin.H{"farms": farms})
}

// GET /farms/:id
// Find a farms
func GetFarm(c *gin.Context) {
	// Get model if exist
	var farm models.Farm
	if err := models.DB.Where("id = ?", c.Param("id")).Preload("FarmAddress").First(&farm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "farm not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"farm": farm})
}
