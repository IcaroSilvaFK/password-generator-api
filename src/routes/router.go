package routes

import (
	"github.com/IcaroSilvaFK/password-generator-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func ApplicationRoutes(router *gin.RouterGroup) {

	router.GET("/passwords/all", controllers.FindAllPass)
	router.POST("/passwords/create", controllers.CreateNewPass)
	router.GET("/passwords/batch", controllers.CreateInBatchPass)

}
