package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	log.Fatal(app.Listen(":3000")) // Start the server on port 3000
	log.Print("Server is running on http://localhost:3000")
}
