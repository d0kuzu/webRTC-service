package rest

import (
	"aisale/api/infrastructure/controllers/phone_controllers"
	"github.com/gin-gonic/gin"
)

func PhoneRoutes(router *gin.Engine) {
	productGroup := router.Group("/phones")
	{
		productGroup.POST("/", phone_controllers.Create)

		productGroup.GET("/:id", phone_controllers.GetById)
		productGroup.PUT("/:id", phone_controllers.UpdateById)
		productGroup.DELETE("/:id", phone_controllers.DeleteById)

		productGroup.GET("/all", phone_controllers.GetAll)

		productGroup.GET("/find/:param", phone_controllers.FindPhone)
	}
}
