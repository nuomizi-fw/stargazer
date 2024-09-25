package service

import "github.com/nuomizi-fw/stargazer/core"

type AuthService interface {
	Register() error
	Login() error
	MFAGenerate() error
	MFAVerify() error
	ForgotPassword() error
}

type authService struct {
	db     core.StargazerDB
	logger core.StargazerLogger
}

func NewAuthService(
	db core.StargazerDB,
	logger core.StargazerLogger,
) AuthService {
	return &authService{db, logger}
}

func (as *authService) Register() error {
	return nil
}

func (as *authService) Login() error {
	return nil
}

func (as *authService) MFAGenerate() error {
	return nil
}

func (as *authService) MFAVerify() error {
	return nil
}

func (as *authService) ForgotPassword() error {
	return nil
}
