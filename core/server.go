package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type StargazerServer struct {
	App *fiber.App
	Api fiber.Router
}

func NewStargazerServer() *StargazerServer {
	app := fiber.New(fiber.Config{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Stargazer",
		AppName:           "Stargazer",
		EnablePrintRoutes: true,
	})

	app.Use(healthcheck.New())
	app.Use(etag.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(pprof.New())
	app.Use(cors.New(
		cors.Config{
			MaxAge: 1728000,
		},
	))

	return &StargazerServer{
		App: app,
		Api: app.Group("/api"),
	}
}

func (s *StargazerServer) Start() error {
	return s.App.Listen(":3000")
}

func (s *StargazerServer) Shutdown() error {
	return s.App.Shutdown()
}
