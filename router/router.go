package router

import (
	"bytes"
	"crypto/ecdsa"
	_ "embed"
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/core"
	sjwt "github.com/nuomizi-fw/stargazer/pkg/jwt"
	"github.com/nuomizi-fw/stargazer/pkg/keystore"
	"github.com/nuomizi-fw/stargazer/service"
	middleware "github.com/oapi-codegen/fiber-middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	Module = fx.Module(
		"router",
		fx.Provide(
			NewStargazerRouter,
		),
	)

	_ api.ServerInterface = (*StargazerRouter)(nil)

	ErrInvalidTokenType = errors.New("invalid token type for this endpoint")
	ErrTokenExpired     = errors.New("token has expired")
)

type StargazerRouter struct {
	logger core.StargazerLogger
	service.StargazerService
}

func isPublicPath(method []byte, path string) bool {
	// Public GET endpoints
	if bytes.Equal(method, []byte("GET")) {
		switch path {
		case "/favicon.ico", sjwt.JWKSPath:
			return true
		}
	}

	// Public POST endpoints
	if bytes.Equal(method, []byte("POST")) {
		switch path {
		case "/api/user/login",
			"/api/user/register":
			return true
		}
	}

	return false
}

func NewStargazerRouter(
	logger core.StargazerLogger,
	config core.StargazerConfig,
	server core.StargazerServer,
	service service.StargazerService,
	ks *keystore.KeyStore,
) StargazerRouter {
	router := StargazerRouter{
		StargazerService: service,
	}

	_, publicKey := ks.GetKeyPair()

	server.App.Get(sjwt.JWKSPath, func(c *fiber.Ctx) error {
		// Generate and serve JWKS
		return c.JSON(sjwt.GenerateJwksJSON(publicKey))
	})

	jwtMiddleware := fiber.Handler(func(c *fiber.Ctx) error {
		if isPublicPath(c.Request().Header.Method(), c.Path()) {
			return c.Next()
		}

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "missing authorization header")
		}

		// Validate the token
		valid, claims, err := sjwt.Validate(tokenString, func() (*ecdsa.PublicKey, error) {
			return publicKey, nil
		})
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Token validation failed: "+err.Error())
		}
		if !valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		// Set user information in request headers
		c.Request().Header.Set("uid", strconv.Itoa(claims.ID))
		c.Request().Header.Set("uname", claims.Username)

		return c.Next()
	})

	swagger, err := api.GetSwagger()
	if err != nil {
		logger.Panic("Failed to get swagger: %s", zap.Error(err))
	}

	u, err := url.Parse(swagger.Servers[0].URL)
	if err != nil {
		panic(err)
	}

	validatorMiddleware := middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	})

	// Register OpenAPI handlers
	api.RegisterHandlersWithOptions(server.App, router, api.FiberServerOptions{
		BaseURL: strings.TrimRight(u.Path, "/"),
		Middlewares: []api.MiddlewareFunc{
			validatorMiddleware, // Run OpenAPI validation first
			jwtMiddleware,       // Then JWT validation
		},
	})

	return router
}
