package model

import (
	"time"

	"github.com/nuomizi-fw/stargazer/core"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("model", fx.Provide(NewStargazerModels))

type BaseModel struct {
	gorm.Model
	ID        int            `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

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
		},
	}
}
