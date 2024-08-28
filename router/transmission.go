package router

import "github.com/nuomizi-fw/stargazer/core"

type TransmissionRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewTransmissionRouter(stargazer core.StargazerServer, logger core.StargazerLogger) TransmissionRouter {
	return TransmissionRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (tr TransmissionRouter) InitRouter() {
	tr.logger.Info("Initializing transmission router")
}
