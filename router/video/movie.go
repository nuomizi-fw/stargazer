package video

import "github.com/nuomizi-fw/stargazer/core"

type MovieRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewMovieRouter(stargazer core.StargazerServer, logger core.StargazerLogger) MovieRouter {
	return MovieRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (sr MovieRouter) InitRouter() {
	sr.logger.Info("Initializing movie router")
}
