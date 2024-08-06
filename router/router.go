package router

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"router",
	fx.Provide(
		NewStargazerRouter,
	),
	// Add new router below
	fx.Provide(
		NewPingRouter,
	),
)

type StargazerRouter interface {
	InitRouter()
}

type StargazerRouters []StargazerRouter

func (sr StargazerRouters) InitRouter() {
	for _, router := range sr {
		router.InitRouter()
	}
}

func NewStargazerRouter(
	pingRouter PingRouter,
) StargazerRouters {
	return StargazerRouters{
		pingRouter,
	}
}
