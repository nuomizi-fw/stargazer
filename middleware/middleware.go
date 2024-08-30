package middleware

import "go.uber.org/fx"

var Module = fx.Module(
	"middleware",
	fx.Options(
		fx.Provide(NewMiddleware),
		// Add new middleware below
		fx.Provide(NewJWTMiddleware),
		fx.Provide(NewCorsMiddleware),
		fx.Provide(NewSwaggerMiddleware),
	),
)

type StargazerMiddleware interface {
	InitMiddleware()
}

type StargazerMiddlewares []StargazerMiddleware

func (sm StargazerMiddlewares) InitMiddleware() {
	for _, middleware := range sm {
		middleware.InitMiddleware()
	}
}

func NewMiddleware() StargazerMiddlewares {
	return StargazerMiddlewares{}
}
