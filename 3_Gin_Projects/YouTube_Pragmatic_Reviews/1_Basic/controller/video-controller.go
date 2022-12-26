package controller

import (
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/1_Basic/entity"
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/1_Basic/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
