package rest

import (
	"aisale/api/infrastructure/controllers/ringostat_controllers"
	"github.com/gin-gonic/gin"
)

func RingostatRoutes(router *gin.Engine) {
	productGroup := router.Group("/ring")
	{
		productGroup.POST("/", ringostat_controllers.Call)
		productGroup.POST("/test", ringostat_controllers.Test)
		productGroup.GET("/signal", ringostat_controllers.Signal)
	}
}
