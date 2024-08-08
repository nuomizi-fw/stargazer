package core

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type StargazerDB struct {
	DB *gorm.DB
}

func (db StargazerDB) AutoMigrate() {
}

func NewStargazerDB(config StargazerConfig, logger StargazerLogger) StargazerDB {
	db, err := gorm.Open(sqlite.Open(config.Database.DBFile), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})
	if err != nil {
		logger.Error("Failed to connect to database", err)
	}

	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(1)
	c.SetConnMaxIdleTime(time.Second * 1000)

	return StargazerDB{DB: db}
}
