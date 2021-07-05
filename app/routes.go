package app

import "github.com/gofiber/fiber/v2"

func _IndexHandler(c *fiber.Ctx) error {
	groups := fetchServices()
	return c.Render("templates/index", &fiber.Map{
		"Title":  config.SiteTitle,
		"Groups": groups,
	})
}

func _HealthHandler(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"healthy": true})
}

func Register(app *fiber.App) {
	app.Get("/", _IndexHandler)
	app.Get("/health", _HealthHandler)
}
