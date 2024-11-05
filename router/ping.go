package router

import "github.com/gofiber/fiber/v2"

func (StargazerRouter) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("Pong!")
}
