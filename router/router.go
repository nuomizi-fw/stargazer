package router

import (
	"github.com/nuomizi-fw/stargazer/oapi"
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

func NewStargazerRouter() StargazerRouter {
	return StargazerRouter{}
}
