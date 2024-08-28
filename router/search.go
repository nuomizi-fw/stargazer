package router

import "github.com/nuomizi-fw/stargazer/core"

type SearchRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewSearchRouter(stargazer core.StargazerServer, logger core.StargazerLogger) SearchRouter {
	return SearchRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (sr SearchRouter) InitRouter() {
	sr.logger.Info("Initializing search router")
}
