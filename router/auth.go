package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/oapi"
	"github.com/samber/lo"
)

func (sr StargazerRouter) Login(ctx *fiber.Ctx) error {
	var loginRequest oapi.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(oapi.ResponseBadRequest{
				Message: lo.ToPtr(err.Error()),
			})
	}

	data, err := sr.Auth.Login(loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(oapi.ResponseInternalServerError{
				Message: lo.ToPtr(err.Error()),
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(oapi.GetLoginResponseOK{
			User: &data,
		})
}

func (sr StargazerRouter) Register(ctx *fiber.Ctx) error {
	var registerRequest oapi.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return err
	}

	data, err := sr.Auth.Register(registerRequest)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&data)
}

func (sr StargazerRouter) Refresh(ctx *fiber.Ctx) error {
	var refreshRequest oapi.RefreshTokenRequest

	if err := ctx.BodyParser(&refreshRequest); err != nil {
		return err
	}

	data, err := sr.Auth.RefreshToken(refreshRequest)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(oapi.PostRefreshTokenResponseOK{
		AccessToken:  lo.ToPtr(data.AccessToken),
		RefreshToken: lo.ToPtr(data.RefreshToken),
		ExpiresIn:    lo.ToPtr(data.ExpiresIn),
	})
}
