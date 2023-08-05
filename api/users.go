package api

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest struct {
	Email string `json:"email"`
}

func (api *AppApi) CreateUser(ctx *fiber.Ctx) error {
	var req CreateUserRequest

	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	save, err := api.client.User.Create().SetEmail(req.Email).Save(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(save)
}

func (api *AppApi) GetAllUsers(ctx *fiber.Ctx) error {
	all, err := api.client.User.Query().All(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(all)
}
