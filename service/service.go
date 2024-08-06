package service

import "go.uber.org/fx"

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewService),
		// Add new service below
		fx.Provide(
			NewPingService,
		),
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
