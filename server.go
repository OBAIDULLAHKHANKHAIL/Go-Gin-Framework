package main

import (
	"golang-gin-poc/controller"
	"golang-gin-poc/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message":"hi this is my code",
	// 	})
	// })

	server.GET("/videos", func(ctx *gin.Context){
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context){
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
