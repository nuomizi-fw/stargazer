package router

import "github.com/nuomizi-fw/stargazer/core"

type MangaRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewMangaRouter(stargazer core.StargazerServer, logger core.StargazerLogger) MangaRouter {
	return MangaRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (mr MangaRouter) InitRouter() {
	mr.logger.Info("Initializing manga router")
}
