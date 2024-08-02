package service

type PingService interface {
	Ping() string
}

type pingService struct{}

func (s *pingService) Ping() string {
	return "pong"
}

func NewPingService() PingService {
	return &pingService{}
}
