package router

import "github.com/nuomizi-fw/stargazer/core"

type AuthRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewAuthRouter(stargazer core.StargazerServer, logger core.StargazerLogger) AuthRouter {
	return AuthRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (ar AuthRouter) InitRouter() {
	ar.logger.Info("Initializing auth router")
}
