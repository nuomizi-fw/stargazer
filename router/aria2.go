package router

import (
	"github.com/nuomizi-fw/stargazer/core"
)

type Aria2Router struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewAria2Router(stargazer core.StargazerServer, logger core.StargazerLogger) Aria2Router {
	return Aria2Router{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (ar Aria2Router) InitRouter() {
	ar.logger.Info("Initializing aria2 router")
}
