package core

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type StargazerServer struct {
	App *fiber.App
	Api fiber.Router
}

func NewStargazerServer(config StargazerConfig, sl StargazerLogger) StargazerServer {
	app := fiber.New(fiber.Config{
		Prefork:               !config.Server.Debug,
		CaseSensitive:         true,
		StrictRouting:         true,
		ServerHeader:          "Stargazer",
		AppName:               "Stargazer",
		EnablePrintRoutes:     config.Server.Debug,
		DisableStartupMessage: !config.Server.Debug,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	})

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: sl.Desugar(),
	}))
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(pprof.New())

	return StargazerServer{
		App: app,
	}
}
