package handler

import (
	"net/http"
	"projects/url-shortener/app/handler/model"
	"projects/url-shortener/app/service"

	"github.com/gin-gonic/gin"
)

func CreateShortenedURL(ctx *gin.Context) {
	request := model.Request{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if _, err := service.ShortenURL(&request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}
