package service

import (
	"github.com/nuomizi-fw/stargazer/model"
	"github.com/nuomizi-fw/stargazer/oapi"
	"github.com/nuomizi-fw/stargazer/repository"
)

type AuthService interface {
	Register(register oapi.RegisterRequest) (oapi.User, error)
	Login(login oapi.LoginRequest) (oapi.User, error)
	RefreshToken(refresh oapi.RefreshTokenRequest) (model.RefreshToken, error)
}

type authService struct {
	repo repository.Repository
}

func (a *authService) Register(register oapi.RegisterRequest) (oapi.User, error) {
	return oapi.User{}, nil
}

func (a *authService) Login(login oapi.LoginRequest) (oapi.User, error) {
	return oapi.User{}, nil
}

func (a *authService) RefreshToken(refresh oapi.RefreshTokenRequest) (model.RefreshToken, error) {
	return model.RefreshToken{}, nil
}

func NewAuthService(
	repo repository.Repository,
) AuthService {
	return &authService{
		repo: repo,
	}
}
