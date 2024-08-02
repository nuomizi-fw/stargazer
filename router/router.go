package router

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewStargazerApi),
)

type StargazerRoute interface{}

type StargazerApi struct{}

func NewStargazerApi() *StargazerApi {
	return &StargazerApi{}
}
