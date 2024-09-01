package router

import "github.com/nuomizi-fw/stargazer/core"

type BangumiRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewBangumiRouter(stargazer core.StargazerServer, logger core.StargazerLogger) BangumiRouter {
	return BangumiRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (br BangumiRouter) InitRouter() {
	br.logger.Info("Initializing bangumi router")
}
