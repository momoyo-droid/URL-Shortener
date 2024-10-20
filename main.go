package main

import (
	"projects/url-shortener/app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	rq := gin.Default()

	rq.POST("/shorten", func(ctx *gin.Context) {
		handler.CreateShortenedURL(ctx)
	})
}