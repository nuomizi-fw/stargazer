package middleware

import "github.com/nuomizi-fw/stargazer/core"

type CorsMiddleware struct {
	config core.StargazerConfig
	logger core.StargazerLogger
}

func NewCorsMiddleware(config core.StargazerConfig, logger core.StargazerLogger) CorsMiddleware {
	return CorsMiddleware{
		config: config,
		logger: logger,
	}
}

func (cm CorsMiddleware) InitMiddleware() {
	cm.logger.Info("Cors middleware initialized")
}
