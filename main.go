package main

import (
	"fmt"
	"os"
	"time"

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
	// if err := app.Listen(fmt.Sprintf(":%s", env.ServicePort)); err != nil {
	// 	fmt.Println(err)
	// 	time.Sleep(5 * time.Second)
	// 	os.Exit(1)
	// }
	if err := app.Listen(":5000"); err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}
