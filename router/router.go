package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewStargazerRouter),
)

type StargazerRouter struct{
	
}

func NewStargazerRouter() *StargazerRouter {
	return &StargazerRouter{}
}

func (r *StargazerRouter) Echo(ctx *fiber.Ctx) {
	ctx.SendString("Hello, World!")
}
