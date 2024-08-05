package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewService),
	fx.Provide(
		// Add new service here
		NewPingService,
	),
)

type StargazerService interface {
	Ping() PingService
}

type stargazerService struct {
	ping PingService
}

func (ss *stargazerService) Ping() PingService {
	return ss.ping
}

func NewService() StargazerService {
	return &stargazerService{
		ping: NewPingService(),
	}
}
