//go:generate bash -c "go generate ./ent"
//go:generate bash -c "mkdir -p oapi && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi.yaml api/openapi.yaml"
package main

import (
	"context"
	_ "embed"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/oapi"
	"github.com/nuomizi-fw/stargazer/router"
	"github.com/nuomizi-fw/stargazer/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var (
	//go:embed api/openapi.yaml
	docYAML string
	//go:embed api/index.html
	docHTML string
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

			swagger, err := oapi.GetSwagger()
			if err != nil {
				logger.Panic("Failed to get swagger: %s", zap.Error(err))
			}
			server.App.Use(adaptor.HTTPHandler(NewDocsRouter(swagger, docHTML, docYAML)))

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

func NewDocsRouter(swagger *openapi3.T, docHTML, docYAML string) http.Handler {
	u, err := url.Parse(swagger.Servers[0].URL)
	if err != nil {
		return nil
	}

	apiPath := strings.TrimRight(u.Path, "/")
	docsPath := "/docs" + apiPath

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == docsPath {
			if _, err := w.Write([]byte(docHTML)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		if r.URL.Path == "/openapi.yaml" {
			if _, err := w.Write([]byte(docYAML)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}
