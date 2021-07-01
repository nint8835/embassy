package app

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func Start() error {
	templateEngine := html.NewFileSystem(http.FS(templateFS), ".html.tmpl")
	app := fiber.New(
		fiber.Config{
			Views: templateEngine,
		},
	)

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(staticFS),
		PathPrefix: "static",
	}))

	Register(app)

	app.Listen(":3000")
	return nil
}
