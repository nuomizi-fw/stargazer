package music

import "github.com/nuomizi-fw/stargazer/core"

type MusicRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewMusicRouter(stargazer core.StargazerServer, logger core.StargazerLogger) MusicRouter {
	return MusicRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (mr MusicRouter) InitRouter() {
	mr.logger.Info("Initializing music router")
}
