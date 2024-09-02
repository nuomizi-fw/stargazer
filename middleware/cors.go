package middleware

import "github.com/nuomizi-fw/stargazer/core"

type CorsMiddleware struct {
	config    core.StargazerConfig
	logger    core.StargazerLogger
	stargazer core.StargazerServer
}

func NewCorsMiddleware(
	config core.StargazerConfig,
	logger core.StargazerLogger,
	stargazer core.StargazerServer,
) CorsMiddleware {
	return CorsMiddleware{config, logger, stargazer}
}

func (cm CorsMiddleware) InitMiddleware() {
	cm.logger.Info("Initializing Cors middleware")

	if !cm.config.Server.Cors.Enabled {
		cm.logger.Info("Cors middleware disabled")
		return
	}
}
