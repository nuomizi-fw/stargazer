package router

import "github.com/nuomizi-fw/stargazer/core"

type RssRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewRssRouter(stargazer core.StargazerServer, logger core.StargazerLogger) RssRouter {
	return RssRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (rr RssRouter) InitRouter() {
	rr.logger.Info("Initializing Rss router")
}
