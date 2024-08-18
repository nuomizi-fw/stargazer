package service

type VideoService interface{}

type videoService struct{}

func NewVideoService() VideoService {
	return videoService{}
}
