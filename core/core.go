package core

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewStargazerConfig,
		NewStargazerDB,
		NewStargazerServer,
	),
)
