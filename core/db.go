package core

import "gorm.io/gorm"

type StargazerDB struct {
	DB *gorm.DB
}

func NewStargazerDB() StargazerDB {
	return StargazerDB{}
}
