package core

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
)

type StargazerServer struct {
	App *fiber.App
}

func NewStargazerServer() *StargazerServer {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Stargazer",
		AppName:       "Stargazer",
	})

	app.Use(cors.New(
		cors.Config{
			MaxAge: 1728000,
		},
	))

	return &StargazerServer{
		App: app,
	}
}

func (s *StargazerServer) Start() error {
	return s.App.Listen(":3000")
}

func (s *StargazerServer) Shutdown() error {
	return s.App.Shutdown()
}

func Stargazer(
	lc fx.Lifecycle,
	server *StargazerServer,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Start(); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(); err != nil {
				return err
			}
			return nil
		},
	})
}
