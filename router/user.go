package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/core"
)

type UserRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewUserRouter(stargazer core.StargazerServer, logger core.StargazerLogger) UserRouter {
	return UserRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (ur UserRouter) InitRouter() {
	ur.logger.Info("Initializing user router")
}

func (ur *UserRouter) GetUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) GetUsers(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) CreateUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) DeleteUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) SetUserRole(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) SetUserPermission(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) ResetPassword(ctx *fiber.Ctx) error {
	return nil
}
