package router

import (
	"github.com/gofiber/fiber/v2"
)

func (pa *StargazerApi) GetPing(ctx *fiber.Ctx) error {
	return ctx.SendString(pa.PingService.GetPing())
}
