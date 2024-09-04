package service

type MusicService interface{}

type musicService struct{}

func NewMusicService() MusicService {
	return &musicService{}
}
