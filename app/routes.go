package app

import "github.com/gofiber/fiber/v2"

func _IndexHandler(c *fiber.Ctx) error {
	return c.Render("templates/index", &fiber.Map{
		"Title":  "Embassy",
		"Groups": []fiber.Map{},
	})
}

func Register(app *fiber.App) {
	app.Get("/", _IndexHandler)
}
