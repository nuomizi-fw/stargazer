package core

import "go.uber.org/fx"

var Module = fx.Module(
	"core",
	fx.Options(
		fx.Provide(
			NewStargazerConfig,
			NewStargazerDB,
			NewStargazerLogger,
			NewStargazerServer,
		),
	),
)
