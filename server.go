package main

import (
	"github.com/ferminhg/golang-gin-poc/infrastructure"
	"github.com/ferminhg/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	adService := service.New()
	adController := infrastructure.New(adService)

	server.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "👍",
		})
	})

	server.GET("/ads", func(context *gin.Context) {
		context.JSON(200, adController.FindAll())
	})

	server.POST("/ads", func(context *gin.Context) {
		context.JSON(200, adController.Save(context))
	})

	server.Run(":8080")
}
