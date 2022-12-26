package main

import (
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/1_Basic/controller"
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/1_Basic/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, VideoController.Save(ctx))
	})

	server.Run(":8080")
}

// go mod init home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/1_Basic
// export GO111MODULE="on"
// go get github.com/gin-gonic/gin
