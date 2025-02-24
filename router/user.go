package router

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/service"
	"github.com/samber/lo"
)

func (sr StargazerRouter) Login(ctx *fiber.Ctx) error {
	var loginRequest api.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(api.ResponseBadRequest{
				Message: lo.ToPtr(err.Error()),
			})
	}

	data, err := sr.User.Login(loginRequest)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			return ctx.Status(fiber.StatusUnauthorized).
				JSON(api.ResponseUnauthorized{
					Message: lo.ToPtr("Invalid username or password"),
				})
		default:
			return ctx.Status(fiber.StatusInternalServerError).
				JSON(api.ResponseInternalServerError{
					Message: lo.ToPtr("Internal server error"),
				})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(api.GetLoginResponseOK{
		Token: data,
	})
}

func (sr StargazerRouter) Register(ctx *fiber.Ctx) error {
	var registerRequest api.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(api.ResponseBadRequest{
				Message: lo.ToPtr(err.Error()),
			})
	}

	data, err := sr.User.Register(registerRequest)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserExists):
			return ctx.Status(fiber.StatusConflict).
				JSON(api.ResponseConflictError{
					Message: lo.ToPtr("User already exists"),
				})
		case errors.Is(err, service.ErrInvalidUsername):
			return ctx.Status(fiber.StatusBadRequest).
				JSON(api.ResponseBadRequest{
					Message: lo.ToPtr("Invalid username format"),
				})
		case errors.Is(err, service.ErrInvalidEmail):
			return ctx.Status(fiber.StatusBadRequest).
				JSON(api.ResponseBadRequest{
					Message: lo.ToPtr("Invalid email format"),
				})
		case errors.Is(err, service.ErrWeakPassword):
			return ctx.Status(fiber.StatusBadRequest).
				JSON(api.ResponseBadRequest{
					Message: lo.ToPtr("Password is too weak"),
				})
		default:
			return ctx.Status(fiber.StatusInternalServerError).
				JSON(api.ResponseInternalServerError{
					Message: lo.ToPtr("Internal server error"),
				})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (sr StargazerRouter) Refresh(ctx *fiber.Ctx) error {
	var refreshRequest api.RefreshTokenRequest

	if err := ctx.BodyParser(&refreshRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(api.ResponseBadRequest{
				Message: lo.ToPtr(err.Error()),
			})
	}

	data, err := sr.User.RefreshToken(refreshRequest)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidToken):
			return ctx.Status(fiber.StatusUnauthorized).
				JSON(api.ResponseUnauthorized{
					Message: lo.ToPtr("Invalid or expired refresh token"),
				})
		default:
			return ctx.Status(fiber.StatusInternalServerError).
				JSON(api.ResponseInternalServerError{
					Message: lo.ToPtr("Internal server error"),
				})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (sr StargazerRouter) GetUserProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("Pong!")
}
