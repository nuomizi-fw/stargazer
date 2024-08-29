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
		auth.Post("/mfa/generate", ar.MFAGenerate)
		auth.Post("/mfa/login", ar.MFALogin)
		auth.Post("/forgot-password", ar.ForgotPassword)
	}
}

func (as *AuthRouter) Register(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) Login(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) MFAGenerate(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) MFALogin(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) ForgotPassword(ctx *fiber.Ctx) error {
	return nil
}
