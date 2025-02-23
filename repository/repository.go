package repository

import (
	"context"

	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/ent/user"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"go.uber.org/fx"
)

var Module = fx.Module("repository", fx.Provide(NewStargazerRepository))

type Repository interface {
	// User management
	UserExists(username string) (bool, error)
	CreateUser(user api.User, hashedPassword string) error
	GetUserWithPassword(username string) (api.User, string, error)
}

type StargazerRepository struct {
	db core.StargazerDB
}

func (r *StargazerRepository) UserExists(username string) (bool, error) {
	count, err := r.db.User.Query().
		Where(user.UsernameEQ(username)).
		Count(context.Background())
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *StargazerRepository) CreateUser(user api.User, hashedPassword string) error {
	_, err := r.db.User.Create().
		SetUsername(user.Username).
		SetEmail(string(user.Email)).
		SetPassword(hashedPassword).
		Save(context.Background())

	return err
}

func (r *StargazerRepository) GetUserWithPassword(username string) (api.User, string, error) {
	u, err := r.db.User.Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	if err != nil {
		return api.User{}, "", err
	}

	email := openapi_types.Email(u.Email)
	return api.User{
		Username: u.Username,
		Email:    email,
	}, u.Password, nil
}

func NewStargazerRepository(
	db core.StargazerDB,
) Repository {
	return &StargazerRepository{
		db: db,
	}
}
