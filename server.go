package main

import (
	"github.com/ferminhg/golang-gin-poc/infrastructure"
	"github.com/ferminhg/golang-gin-poc/infrastructure/middleware"
	"github.com/ferminhg/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"os"
)

func serverLogOutput() {
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	serverLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		gindump.Dump(),
		middleware.BasicAuth(),
	)

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
