package main

import (
	_ "embed"
	"strings"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	// app.Use(limiter.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:               200,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(compress.New())
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-type", "text/html; charset=utf-8")
		return c.SendString(strings.ReplaceAll(indexTemplate, "{{.Insult}}", InsultsList.GetRandomInsult()))
	})
	rg := app.Group("/api/v1")
	v1 := rg.Group("/insults")
	v1.Get("/random", func(c *fiber.Ctx) error {
		return c.SendString(InsultsList.GetRandomInsult())
	})
	v1.Get("/endusers", func(c *fiber.Ctx) error {
		return c.JSON(EnduserInspiredInsults)
	})
	v1.Get("/endusers/random", func(c *fiber.Ctx) error {
		return c.SendString(EnduserInspiredInsults.GetRandomInsult())
	})
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(InsultsList)
	})
	app.Listen(":8080")
}
