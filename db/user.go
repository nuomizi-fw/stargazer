package db

import (
	"context"

	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/ent/user"
	"github.com/oapi-codegen/runtime/types"
)

func UserExists(username string) (bool, error) {
	count, err := db.User.Query().
		Where(user.UsernameEQ(username)).
		Count(context.Background())
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CreateUser(user api.User, hashedPassword string) error {
	_, err := db.User.Create().
		SetUsername(user.Username).
		SetEmail(string(user.Email)).
		SetPassword(hashedPassword).
		Save(context.Background())

	return err
}

func GetUserWithPassword(username string) (api.User, string, error) {
	u, err := db.User.Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	if err != nil {
		return api.User{}, "", err
	}

	email := types.Email(u.Email)
	return api.User{
		Username: u.Username,
		Email:    email,
	}, u.Password, nil
}
