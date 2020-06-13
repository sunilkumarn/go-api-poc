package server

import (
	"leads_module/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/leads", controllers.GetLeads)
	r.GET("/leads/:id", controllers.GetLead)

	return r
}
