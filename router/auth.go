package router

import (
	"github.com/gofiber/fiber/v2"
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
		auth.Post("/register", ar.Register)
		auth.Post("/login", ar.Login)
		auth.Post("/mfa-login", ar.MFALogin)
		auth.Post("/reset-password", ar.ResetPassword)
	}
}

func (ar AuthRouter) Register(ctx *fiber.Ctx) error {
	return nil
}

func (ar AuthRouter) Login(ctx *fiber.Ctx) error {
	return nil
}

func (ar AuthRouter) MFALogin(ctx *fiber.Ctx) error {
	return nil
}

func (ar AuthRouter) ResetPassword(ctx *fiber.Ctx) error {
	return nil
}
