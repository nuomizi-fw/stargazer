package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewServices),
)

type Services struct{}

func NewServices() *Services {
	return &Services{}
}
