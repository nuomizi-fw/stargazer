package router

import "github.com/nuomizi-fw/stargazer/core"

type VideoRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewVideoRouter(stargazer core.StargazerServer, logger core.StargazerLogger) VideoRouter {
	return VideoRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (vr VideoRouter) InitRouter() {
	vr.logger.Info("Initializing video router")
}
