package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"

	"github.com/muhammadluth/log"
)

func main() {
	env := serviceEnv()

	log.SetupLogging(env.LogPath)

	app := fiber.New()
	app.Use(etag.New())
	app.Use(compress.New())

	// HEALTH CHECK
	app.Get("/", doMiddleware(), func(ctx *fiber.Ctx) error {
		uniqueID := ctx.Locals("UniqueID").(string)
		return ctx.JSON(fiber.Map{
			"id":      uniqueID,
			"message": "Hello, Welcome to My Apps!",
		})
	})

	fmt.Printf("Listening on port : %s \n", env.ServicePort)
	log.Event(fmt.Sprintf("Listening on port : %s \n", env.ServicePort))
	fmt.Println("Ready to serve")
	if err := app.Listen(fmt.Sprintf(":%s", env.ServicePort)); err != nil {
		log.Fatal(err)
	}
}

func doMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		randomHex, _ := doRandomHex(10)
		uniqueID := strings.ToUpper(randomHex)

		log.Message(
			uniqueID,
			"IN",
			"CLIENT",
			"",
			"GO-FIBER SERVICE",
			"INCOMING REQUEST",
			string(ctx.OriginalURL()),
			"",
			string(ctx.Request().Body()),
		)

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		ctx.Locals("UniqueID", uniqueID)

		log.Message(
			uniqueID,
			"OUT",
			"CLIENT",
			"",
			"GO-FIBER SERVICE",
			"OUTGOING RESPONSE",
			string(ctx.OriginalURL()),
			"",
			string(ctx.Response().Body()),
		)
		return ctx.Next()
	}
}

func doRandomHex(number int) (string, error) {
	bytes := make([]byte, number)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
