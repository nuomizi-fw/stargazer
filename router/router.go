package router

import (
	"bytes"
	"strconv"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/oapi"
	sjwt "github.com/nuomizi-fw/stargazer/pkg/jwt"
	"github.com/nuomizi-fw/stargazer/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"router",
		fx.Provide(
			NewStargazerRouter,
		),
	)

	_ oapi.ServerInterface = (*StargazerRouter)(nil)
)

type StargazerRouter struct{}

func NewStargazerRouter(
	config core.StargazerConfig,
	server core.StargazerServer,
	service service.StargazerService,
) StargazerRouter {
	router := StargazerRouter{}

	server.App.Use(jwtware.New(jwtware.Config{
		// Skip middleware for localhost connections
		Filter: func(c *fiber.Ctx) bool {
			if c.IP() == "::1" || c.IP() == "127.0.0.1" {
				return true
			}

			if bytes.Equal(c.Request().Header.Method(), []byte("GET")) && c.Path() == "/ping" {
				return true
			}

			if bytes.Equal(c.Request().Header.Method(), []byte("POST")) && (c.Path() == "/login" || c.Path() == "/register") {
				return true
			}

			return false
		},
		// Use secret key directly
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.Server.Secret),
		},
		// Success handler
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			if claims, ok := token.Claims.(*sjwt.Claims); ok && token.Valid {
				c.Request().Header.Set("user_id", strconv.Itoa(claims.ID))
				c.Request().Header.Set("username", claims.Username)
				return c.Next()
			}
			return fiber.ErrUnauthorized
		},
		// Error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		},
		// Token lookup in header and query
		TokenLookup: "header:Authorization,query:token",
		// Remove "Bearer " prefix
		AuthScheme: "Bearer",
	}))

	oapi.RegisterHandlers(server.App, router)

	return router
}
