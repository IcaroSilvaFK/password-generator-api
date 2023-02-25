package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/password-generator-api/src/services"
	"github.com/gin-gonic/gin"
)

func FindAllPass(ctx *gin.Context) {

	passRepository := services.NewPassService()

	ctx.JSON(http.StatusOK, gin.H{
		"passwords": passRepository.FindAllPass(),
	})
}
