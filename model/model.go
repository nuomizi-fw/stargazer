package model

import (
	"github.com/nuomizi-fw/stargazer/core"
	"go.uber.org/fx"
)

var Module = fx.Module("model", fx.Provide(NewStargazerModels))

type StargazerModel struct {
	models []interface{}
	db     core.StargazerDB
	logger core.StargazerLogger
	config core.StargazerConfig
}

func (sm StargazerModel) AutoMigrate() {
	if !sm.config.Database.Migrate {
		sm.logger.Info("Database migration is disabled")
		return
	}

	sm.logger.Info("Auto migrating models")
	if err := sm.db.AutoMigrate(sm.models...); err != nil {
		sm.logger.Error("Failed to auto migrate models", err)
	}
}

func NewStargazerModels(
	db core.StargazerDB,
	logger core.StargazerLogger,
	config core.StargazerConfig,
) StargazerModel {
	return StargazerModel{
		db:     db,
		logger: logger,
		config: config,
		models: []interface{}{
			&User{},
			&Bangumi{},
			&Season{},
			&CastMember{},
		},
	}
}
