package service

import "home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/9_Basic/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

// /home/bjit/Documents/Golang_Practice/3_Gin_Projects/1_Basic/entity

// 1_Basic/entity
//go work use ./entity
