package download

import (
	"github.com/nuomizi-fw/stargazer/core"
)

type BittorrentRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewBittorrentRouter(stargazer core.StargazerServer, logger core.StargazerLogger) BittorrentRouter {
	return BittorrentRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (br BittorrentRouter) InitRouter() {
	br.logger.Info("Initializing bittorrent router")
}
