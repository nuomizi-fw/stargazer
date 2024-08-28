package router

import "github.com/nuomizi-fw/stargazer/core"

type NovelRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewNovelRouter(stargazer core.StargazerServer, logger core.StargazerLogger) NovelRouter {
	return NovelRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (nr NovelRouter) InitRouter() {
	nr.logger.Info("Initializing novel router")
}
