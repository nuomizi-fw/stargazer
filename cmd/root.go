package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nuomizi-fw/stargazer/core"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stargazer",
	Short: "An All-in-One self-hosted solution for your videos, music, manga, novels and more.",
	Long:  "An All-in-One self-hosted solution for your videos, music, manga, novels and more.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stargazer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
