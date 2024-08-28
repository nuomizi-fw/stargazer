package service

import "github.com/gofiber/fiber/v2"

type AuthService interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	MFALogin(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (as *authService) Register(ctx *fiber.Ctx) error {
	return nil
}

func (as *authService) Login(ctx *fiber.Ctx) error {
	return nil
}

func (as *authService) MFALogin(ctx *fiber.Ctx) error {
	return nil
}

func (as *authService) ResetPassword(ctx *fiber.Ctx) error {
	return nil
}
