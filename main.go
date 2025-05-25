package main

import (
	"github.com/gofiber/fiber/v2"
	"BelajarGolang/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/user", handlers.GetUser)

	app.Listen(":3000")
}
