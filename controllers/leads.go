package controllers

import (
	"net/http"

	"leads_module/models"

	"github.com/gin-gonic/gin"
)

// GET /leads
// Get all leads
func GetLeads(c *gin.Context) {
	var leads []models.LeadDetail
	models.DB.Preload("LeadAddress").Preload("LeadSubDetail").Find(&leads)
	c.JSON(http.StatusOK, gin.H{"leads": leads})
}

// GET /leads/:id
// Find a lead
func GetLead(c *gin.Context) {
	// Get model if exist
	var lead models.LeadDetail
	if err := models.DB.Where("id = ?", c.Param("id")).Preload("LeadAddress").Preload("LeadSubDetail").First(&lead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lead not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lead": lead})
}
