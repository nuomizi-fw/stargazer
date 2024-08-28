package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nuomizi-fw/stargazer/core"
)

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTMiddleware struct {
	config core.StargazerConfig
	logger core.StargazerLogger
}

func NewJWTMiddleware(config core.StargazerConfig, logger core.StargazerLogger) JWTMiddleware {
	return JWTMiddleware{
		config: config,
		logger: logger,
	}
}

func (jm JWTMiddleware) InitMiddleware() {
	jm.logger.Info("JWT middleware initialized")
}
