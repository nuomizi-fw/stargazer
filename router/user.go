package router

import "github.com/nuomizi-fw/stargazer/core"

type UserRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewUserRouter(stargazer core.StargazerServer, logger core.StargazerLogger) UserRouter {
	return UserRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (ur UserRouter) InitRouter() {
	ur.logger.Info("Initializing user router")
}
