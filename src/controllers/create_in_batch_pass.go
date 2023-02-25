package controllers

import (
	"net/http"
	"strconv"

	"github.com/IcaroSilvaFK/password-generator-api/src/services"
	"github.com/gin-gonic/gin"
)

func CreateInBatchPass(ctx *gin.Context) {

	param, oki := ctx.GetQuery("limit")

	passService := services.NewPassService()

	if !oki {
		pass, err := passService.RandomPass(10)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"passwords": pass,
		})

		return
	}

	castingFromInt, err := strconv.Atoi(param)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	pass, err := passService.RandomPass(castingFromInt)

	ctx.JSON(http.StatusOK, gin.H{
		"passwords": pass,
	})

}
