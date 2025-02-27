//go:generate bash -c "go generate ./ent"
//go:generate bash -c "go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types,spec,fiber -package api api/openapi.yaml > api/api.gen.go"
package main

import (
	_ "embed"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nuomizi-fw/stargazer/db"
	"github.com/nuomizi-fw/stargazer/pkg/config"
	"github.com/nuomizi-fw/stargazer/pkg/keystore"
	"github.com/nuomizi-fw/stargazer/router"
	"github.com/nuomizi-fw/stargazer/service"
	"go.uber.org/zap"
)

func main() {
	// Initialize components only once
	config := config.NewStargazerConfig() // Initialize config

	file, _ := os.OpenFile(
		config.Logger.LogPath+"/"+config.Logger.LogName+"."+config.Logger.LogExt,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666,
	)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)

	db.NewStargazerDB(config) // Initialize database
	var err error
	keyStore, err := keystore.NewKeyStore() // Initialize keystore
	if err != nil {
		log.Panic("Failed to initialize keystore", zap.Error(err))
	}
	stargazerService := service.NewStargazerService() // Initialize services

	// Create Fiber app
	app := fiber.New(fiber.Config{
		Prefork:               config.Server.Debug,
		CaseSensitive:         true,
		StrictRouting:         true,
		ServerHeader:          "Stargazer",
		AppName:               "Stargazer",
		EnablePrintRoutes:     config.Server.Debug,
		DisableStartupMessage: !config.Server.Debug,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(compress.New())

	// Initialize router
	router.NewStargazerRouter(config, app, stargazerService, keyStore)

	// migrate datebase
	if config.Database.Migrate {
		log.Info("Migrating database...")

		db.AutoMigrate()
	}

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(config.Server.Port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	log.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Info("Running cleanup tasks...")

	// Your cleanup tasks go here
	db.CloseStargazerDB()
	log.Info("Stargazer was successful shutdown.")
}
