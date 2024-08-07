package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/service"
)

type PingRouter struct {
	stargazer   core.StargazerServer
	logger      core.StargazerLogger
	pingService service.PingService
}

func NewPingRouter(stargazer core.StargazerServer, logger core.StargazerLogger, pingService service.PingService) PingRouter {
	return PingRouter{
		stargazer:   stargazer,
		logger:      logger,
		pingService: pingService,
	}
}

func (pr PingRouter) InitRouter() {
	pr.stargazer.Api.Get("/ping", pr.GetPing)
}

func (pr *PingRouter) GetPing(ctx *fiber.Ctx) error {
	return ctx.SendString(pr.pingService.GetPing())
}
