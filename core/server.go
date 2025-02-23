package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type StargazerServer struct {
	App *fiber.App
}

func NewStargazerServer(config StargazerConfig, sl StargazerLogger) StargazerServer {
	app := fiber.New(fiber.Config{
		Prefork: !config.Server.Debug,
		// CaseSensitive:         true,
		// StrictRouting:         true,
		ServerHeader:          "Stargazer",
		AppName:               "Stargazer",
		EnablePrintRoutes:     config.Server.Debug,
		DisableStartupMessage: !config.Server.Debug,
	})

	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))
	app.Use(logger.New())
	app.Use(compress.New())

	return StargazerServer{
		App: app,
	}
}
