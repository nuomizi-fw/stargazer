package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/router"
	"github.com/nuomizi-fw/stargazer/service"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "stargazer",
	Short: "An All-in-One self-hosted solution for your videos, music, manga, novels and more.",
	Long:  "An All-in-One self-hosted solution for your videos, music, manga, novels and more.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		kill := make(chan os.Signal, 1)
		signal.Notify(kill, syscall.SIGTERM, syscall.SIGINT)

		go func() {
			<-kill
			cancel()
		}()

		app := fx.New(
			core.Module,
			router.Module,
			service.Module,
			fx.Invoke(core.Stargazer),
		)

		if err := app.Start(ctx); err != nil {
			if err != context.Canceled {
				zap.L().Fatal("Failed to start application", zap.Error(err))
			}
		}

		<-ctx.Done()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
