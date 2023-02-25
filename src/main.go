package main

import (
	"log"

	"github.com/IcaroSilvaFK/password-generator-api/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to loading .env file create your file or delete and create")
	}
	app := gin.Default()

	app.Use(cors.Default())
	router := app.Group("het/api/_v1")

	routes.ApplicationRoutes(router)

	app.Run()
}
