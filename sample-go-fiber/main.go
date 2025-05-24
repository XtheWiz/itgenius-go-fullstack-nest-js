package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About Page")
	})
	app.Listen(":3000")
}
