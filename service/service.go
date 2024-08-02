package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewService),
)

type StargazerService interface {
	Ping() PingService

	User() UserService
}

type stargazerService struct {
	ping PingService
	user UserService
}

func (ss *stargazerService) Ping() PingService {
	return ss.ping
}

func (ss *stargazerService) User() UserService {
	return ss.user
}

func NewService() StargazerService {
	return &stargazerService{
		ping: NewPingService(),
		user: NewUserService(),
	}
}
