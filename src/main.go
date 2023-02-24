package main

import (
	"github.com/IcaroSilvaFK/password-generator-api/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	app.Use(cors.Default())
	router := app.Group("het/api/_v1")

	routes.ApplicationRoutes(router)

	app.Run()

}
