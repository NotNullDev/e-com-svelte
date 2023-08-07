package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/notnulldev/e-com-svelte/api"
	"github.com/notnulldev/e-com-svelte/ent"
	"log"
	"net/http"
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
	app.Use(cors.New())
	app.Use(recover.New())

	files := app.Group("/static")
	files.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir("./files"),
		MaxAge: 20,
		Browse: true,
	}))

	appApi := api.NewAppApi(client)

	usersApi := app.Group("/users")
	usersApi.Get("/", appApi.GetAllUsers)
	usersApi.Post("/", appApi.CreateUser)

	productsApi := app.Group("/products")
	productsApi.Get("/", appApi.GetAllProducts)
	productsApi.Get("/my", appApi.GetMyProducts)
	productsApi.Get("/:id", appApi.GetProductById)
	productsApi.Post("/", appApi.CreateProduct)
	productsApi.Put("/", appApi.UpdateProduct)

	err = app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
