package router

import (
	"bytes"
	"crypto/ecdsa"
	"strconv"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/oapi"
	sjwt "github.com/nuomizi-fw/stargazer/pkg/jwt"
	"github.com/nuomizi-fw/stargazer/pkg/logger"
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

	_, publicKey := service.User.GetKeyPair()

	jswkJSON, err := sjwt.GenerateJwksJSON(publicKey)
	if err != nil {
		logger.Fatalf("Failed to generate JWKS JSON: %s", err)
	}

	server.App.Use(adaptor.HTTPHandler(sjwt.JWKSHandler(jswkJSON)))

	server.App.Use(jwtware.New(jwtware.Config{
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
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)

			valid, claims, err := sjwt.Validate(
				token.Raw,
				func() (*ecdsa.PublicKey, error) {
					_, publicKey := service.User.GetKeyPair()
					return publicKey, nil
				})
			if err != nil || !valid {
				return fiber.ErrUnauthorized
			}

			c.Request().Header.Set("user_id", strconv.Itoa(claims.ID))
			c.Request().Header.Set("username", claims.Username)

			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		},
		TokenLookup: "header:Authorization,query:token",
		AuthScheme:  "Bearer",
		KeyFunc: func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fiber.ErrUnauthorized
			}

			_, publicKey := service.User.GetKeyPair()
			return publicKey, nil
		},
		JWKSetURLs: []string{sjwt.JWKSPath},
	}))

	oapi.RegisterHandlers(server.App, router)

	return router
}
