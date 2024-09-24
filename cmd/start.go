/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/middleware"
	"github.com/nuomizi-fw/stargazer/model"
	"github.com/nuomizi-fw/stargazer/router"
	"github.com/nuomizi-fw/stargazer/service"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Stargazer server",
	Run: func(cmd *cobra.Command, args []string) {
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
			middleware.Module,
			model.Module,
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
	},
}

func StartStargazer(
	lc fx.Lifecycle,
	config core.StargazerConfig,
	db core.StargazerDB,
	logger core.StargazerLogger,
	server core.StargazerServer,
	middleware middleware.StargazerMiddlewares,
	model model.StargazerModel,
	router router.StargazerRouters,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.InitRouter()
			middleware.InitMiddleware()
			model.AutoMigrate()

			go func() {
				if config.Server.TLS.Enabled {
					if err := server.App.ListenTLS(config.Server.Port, config.Server.TLS.CertFile, config.Server.TLS.KeyFile); err != nil {
						logger.Panic("Failed to start https server: %s", zap.Error(err))
					}
				} else {
					if err := server.App.Listen(config.Server.Port); err != nil {
						logger.Panic("Failed to start http server: %s", zap.Error(err))
					}
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

func init() {
	rootCmd.AddCommand(startCmd)
}
