package main

import (
	_ "embed"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-type", "text/html; charset=utf-8")
		return c.SendString(strings.ReplaceAll(indexTemplate, "{{.Insult}}", InsultsList.GetRandomInsult()))
	})

	app.Listen(":8080")
}
