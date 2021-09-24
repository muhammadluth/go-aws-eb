package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func main() {
	env := serviceEnv()
	app := fiber.New()

	app.Use(etag.New())
	app.Use(compress.New())

	// HEALTH CHECK
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My Apps!"})
	})

	fmt.Println("Listening on port : " + env.ServicePort)
	fmt.Println("Ready to serve")
	app.Listen(fmt.Sprintf(":%s", env.ServicePort))
}
