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
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/core"
	sjwt "github.com/nuomizi-fw/stargazer/pkg/jwt"
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
		case "/api/auth/login",
			"/api/auth/register":
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
) StargazerRouter {
	_, publicKey := service.User.GetKeyPair()

	// Generate and serve JWKS
	jswkJSON, err := sjwt.GenerateJwksJSON(publicKey)
	if err != nil {
		logger.Fatalf("Failed to generate JWKS JSON: %s", err)
	}
	server.App.Get(sjwt.JWKSPath, adaptor.HTTPHandler(sjwt.JWKSHandler(jswkJSON)))

	router := StargazerRouter{
		StargazerService: service,
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		logger.Panic("Failed to get swagger: %s", zap.Error(err))
	}

	u, err := url.Parse(swagger.Servers[0].URL)
	if err != nil {
		panic(err)
	}

	jwtMiddleware := jwtware.New(jwtware.Config{
		Filter: func(c *fiber.Ctx) bool {
			return isPublicPath(c.Request().Header.Method(), c.Path())
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if errors.Is(err, ErrInvalidTokenType) {
				return fiber.NewError(fiber.StatusForbidden, "Invalid token type for this endpoint")
			}
			if errors.Is(err, ErrTokenExpired) {
				return fiber.NewError(fiber.StatusUnauthorized, "Token has expired")
			}
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or missing authentication")
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			if token == nil {
				return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
			}

			// Validate token and extract claims
			valid, claims, err := sjwt.Validate(token.Raw, func() (*ecdsa.PublicKey, error) {
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
			c.Request().Header.Set("client_type", claims.ClientType)

			return c.Next()
		},
		TokenLookup: "header:Authorization,query:token",
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.ES256,
			Key:    publicKey,
		},
	})

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
