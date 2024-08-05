package router

import (
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewStargazerApi),
)

type StargazerApi struct {
	Stargazer *core.StargazerServer

	PingService service.PingService
}

func (sa *StargazerApi) InitStargazerApi() {
	sa.Stargazer.Api.Get("/ping", sa.GetPing)
}

func NewStargazerApi(
	stargazer *core.StargazerServer,
	pingService service.PingService,
) *StargazerApi {
	return &StargazerApi{
		Stargazer:   stargazer,
		PingService: pingService,
	}
}
