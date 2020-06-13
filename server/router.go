package server

import (
	"farm_tales/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/farms", controllers.GetFarms)
	r.GET("/farms/:id", controllers.GetFarm)

	return r
}
