package service

type PingService interface {
	GetPing() string
}

type pingService struct{}

func (s *pingService) GetPing() string {
	return "pong"
}

func NewPingService() PingService {
	return &pingService{}
}
