package router

import (
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/service"
)

type AuthRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
	auth      service.AuthService
}

func NewAuthRouter(
	stargazer core.StargazerServer,
	logger core.StargazerLogger,
	auth service.AuthService,
) AuthRouter {
	return AuthRouter{
		stargazer: stargazer,
		logger:    logger,
		auth:      auth,
	}
}

func (ar AuthRouter) InitRouter() {
	ar.logger.Info("Initializing auth router")

	auth := ar.stargazer.App.Group("/auth")
	{
		auth.Post("/register", ar.auth.Register)
		auth.Post("/login", ar.auth.Login)
		auth.Post("/mfa-login", ar.auth.MFALogin)
		auth.Post("/reset-password", ar.auth.ResetPassword)
	}
}
