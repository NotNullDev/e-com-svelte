package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/notnulldev/e-com-svelte/ent"
	"log"
)

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

		created, err := client.Product.Create().SetName(p.Name).SetPrice(p.Price).Save(ctx.Context())
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
