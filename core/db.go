package core

import (
	"gorm.io/gorm"
)

type StargazerDB struct {
	DB *gorm.DB
}

func NewStargazerDB(config StargazerConfig, logger StargazerLogger) StargazerDB {
	return StargazerDB{}
}
