package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/password-generator-api/src/services"
	"github.com/gin-gonic/gin"
)

type RequestInput struct {
	Pass string `json:"pass"`
}

func CreateNewPass(ctx *gin.Context) {

	var requestBody RequestInput

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	passService := services.NewPassService()

	pass, err := passService.NewPass(requestBody.Pass)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"pass":    pass,
	})

}
