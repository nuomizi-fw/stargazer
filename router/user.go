package router

import "github.com/gofiber/fiber/v2"

func (StargazerRouter) GetUserProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("Pong!")
}
