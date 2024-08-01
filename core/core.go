package core

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewConfig), fx.Provide(NewStargazerServer))
