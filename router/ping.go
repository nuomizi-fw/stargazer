package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/service"
)

type PingRouter struct {
	Stargazer   core.StargazerServer
	PingService service.PingService
}

func NewPingRouter(stargazer core.StargazerServer, pingService service.PingService) PingRouter {
	return PingRouter{
		Stargazer:   stargazer,
		PingService: pingService,
	}
}

func (pr PingRouter) InitRouter() {
	pr.Stargazer.Api.Get("/ping", pr.GetPing)
}

func (pr *PingRouter) GetPing(ctx *fiber.Ctx) error {
	return ctx.SendString(pr.PingService.GetPing())
}
