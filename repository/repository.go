package repository

import (
	"github.com/nuomizi-fw/stargazer/ent"
	"go.uber.org/fx"
)

var Module = fx.Module("repository", fx.Provide(NewStargazerRepository))

type Repository interface{}

type StargazerRepository struct {
	db *ent.Client
}

func NewStargazerRepository(db *ent.Client) *StargazerRepository {
	return &StargazerRepository{db}
}
