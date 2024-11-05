//go:generate bash -c "go generate ./ent"
//go:generate bash -c "mkdir -p oapi && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi.yaml api/openapi.yaml"
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/oapi"
	"github.com/nuomizi-fw/stargazer/router"
	"github.com/nuomizi-fw/stargazer/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-kill
		cancel()
	}()

	app := fx.New(
		fx.WithLogger(func() fxevent.Logger {
			logger := core.NewStargazerLogger()
			return logger.GetFxLogger()
		}),
		core.Module,
		router.Module,
		service.Module,
		fx.Invoke(StartStargazer),
	)

	if err := app.Start(ctx); err != nil {
		if err != context.Canceled {
			log.Fatalf("Failed to start app: %s", err)
		}
	}

	<-ctx.Done()
}

func StartStargazer(
	lc fx.Lifecycle,
	config core.StargazerConfig,
	db core.StargazerDB,
	logger core.StargazerLogger,
	server core.StargazerServer,
	router router.StargazerRouter,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if config.Database.Migrate {
				logger.Info("Migrating database...")

				if err := db.Schema.Create(ctx); err != nil {
					logger.Panic("Failed to migrate database: %s", zap.Error(err))
				}
			}

			oapi.RegisterHandlers(server.App, router)

			go func() {
				if config.Server.TLS.Enabled {
					logger.Fatal(server.App.ListenTLS(config.Server.Port, config.Server.TLS.CertFile, config.Server.TLS.KeyFile))
				} else {
					logger.Fatal(server.App.Listen(config.Server.Port))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.App.Shutdown(); err != nil {
				return err
			}
			return nil
		},
	})
}
