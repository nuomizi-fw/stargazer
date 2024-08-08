package core

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type StargazerServer struct {
	App *fiber.App
	Api fiber.Router
}

func NewStargazerServer(config StargazerConfig, sl StargazerLogger) StargazerServer {
	app := fiber.New(fiber.Config{
		Prefork:           config.Server.Prefork,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Stargazer",
		AppName:           "Stargazer",
		EnablePrintRoutes: true,
	})

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: sl.SugaredLogger.Desugar(),
	}))
	app.Use(recover.New())
	app.Use(compress.New())

	if config.Server.Debug {
		app.Use(pprof.New())
		app.Use(monitor.New())
		app.Use(healthcheck.New())
	}

	apiGroup := app.Group("/api")

	return StargazerServer{
		App: app,
		Api: apiGroup,
	}
}
