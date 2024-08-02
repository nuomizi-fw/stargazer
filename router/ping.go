package router

import "github.com/gofiber/fiber/v2"

func (sa *StargazerApi) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("pong")
}
