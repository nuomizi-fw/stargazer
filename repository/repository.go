package repository

import (
	"context"

	"github.com/nuomizi-fw/stargazer/ent"
	"github.com/nuomizi-fw/stargazer/ent/user"
	"github.com/nuomizi-fw/stargazer/oapi"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

var Module = fx.Module("repository", fx.Provide(NewStargazerRepository))

type Repository interface {
	// User management
	UserExists(username string) (bool, error)
	CreateUser(user oapi.User, hashedPassword string) error
	GetUserWithPassword(username string) (oapi.User, string, error)
}

type StargazerRepository struct {
	db *ent.Client
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

func (r *StargazerRepository) CreateUser(user oapi.User, hashedPassword string) error {
	_, err := r.db.User.Create().
		SetUsername(*user.Username).
		SetEmail(string(*user.Email)).
		SetPassword(hashedPassword).
		Save(context.Background())

	return err
}

func (r *StargazerRepository) GetUserWithPassword(username string) (oapi.User, string, error) {
	u, err := r.db.User.Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	if err != nil {
		return oapi.User{}, "", err
	}

	return oapi.User{
		Username: lo.ToPtr(u.Username),
		Email:    lo.ToPtr(openapi_types.Email(u.Email)),
	}, u.Password, nil
}

func NewStargazerRepository(db *ent.Client) Repository {
	return &StargazerRepository{
		db: db,
	}
}
