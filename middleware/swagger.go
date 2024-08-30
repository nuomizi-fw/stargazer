package middleware

import "github.com/nuomizi-fw/stargazer/core"

type SwaggerMiddleware struct {
	config core.StargazerConfig
	logger core.StargazerLogger
}

func NewSwaggerMiddleware(config core.StargazerConfig, logger core.StargazerLogger) SwaggerMiddleware {
	return SwaggerMiddleware{
		config: config,
		logger: logger,
	}
}

func (cm SwaggerMiddleware) InitMiddleware() {
	cm.logger.Info("Swagger middleware initialized")
}
