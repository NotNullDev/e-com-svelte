package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/notnulldev/e-com-svelte/ent"
	"github.com/notnulldev/e-com-svelte/ent/product"
	"log"
)

type CreateUserRequest struct {
	Email string `json:"email"`
}

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=e-com password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/users", func(ctx *fiber.Ctx) error {
		all, err := client.User.Query().All(ctx.Context())
		if err != nil {
			return err
		}
		return ctx.JSON(all)
	})

	app.Post("/users", func(ctx *fiber.Ctx) error {
		var req CreateUserRequest

		err := ctx.BodyParser(&req)
		if err != nil {
			return err
		}

		save, err := client.User.Create().SetEmail(req.Email).Save(ctx.Context())
		if err != nil {
			return err
		}

		return ctx.JSON(save)
	})

	app.Get("/products/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id", 1)
		if err != nil {
			return err
		}

		all, err := client.Product.Query().Where(product.IDEQ(id)).Only(ctx.Context())
		if err != nil {
			return err
		}
		return ctx.JSON(all)
	})

	app.Get("/products", func(ctx *fiber.Ctx) error {
		all, err := client.Product.Query().All(ctx.Context())
		if err != nil {
			return err
		}
		return ctx.JSON(all)
	})

	app.Post("/products", func(ctx *fiber.Ctx) error {
		var p ent.Product

		err := ctx.BodyParser(&p)
		if err != nil {
			return err
		}

		created, err := client.Product.Create().
			SetName(p.Name).
			SetPrice(p.Price).
			SetPreviewURL(p.PreviewURL).
			SetCategories(p.Categories).
			Save(ctx.Context())
		if err != nil {
			return err
		}

		return ctx.JSON(created)
	})

	err = app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
