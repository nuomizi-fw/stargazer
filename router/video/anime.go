package video

import "github.com/nuomizi-fw/stargazer/core"

type AnimeRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewAnimeRouter(stargazer core.StargazerServer, logger core.StargazerLogger) AnimeRouter {
	return AnimeRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (sr AnimeRouter) InitRouter() {
	sr.logger.Info("Initializing anime router")
}
