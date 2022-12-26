package main

import (
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/3_Basic/controller"
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/3_Basic/middlewares"
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/3_Basic/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
		}
	})

	server.Run(":8080")
}

// go mod init home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/3_Basic
// export GO111MODULE="on"
// go get github.com/gin-gonic/gin
// go get github.com/tpkeeper/gin-dump
