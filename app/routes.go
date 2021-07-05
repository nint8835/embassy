package app

import "github.com/gofiber/fiber/v2"

func _IndexHandler(c *fiber.Ctx) error {
	groups := fetchServices()
	return c.Render("templates/index", &fiber.Map{
		"Title":  config.SiteTitle,
		"Groups": groups,
	})
}

func Register(app *fiber.App) {
	app.Get("/", _IndexHandler)
}
