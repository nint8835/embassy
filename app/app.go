package app

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	consul "github.com/hashicorp/consul/api"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ConsulEndpoint string `default:"127.0.0.1:8500" split_words:"true"`
	TagPrefix      string `default:"embassy" split_words:"true"`
	SiteTitle      string `default:"Embassy" split_words:"true"`
}

//go:embed templates/*
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

var consulClient *consul.Client
var config Config

func Start() error {
	err := envconfig.Process("embassy", &config)
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	consulClient, err = consul.NewClient(&consul.Config{Address: config.ConsulEndpoint})
	if err != nil {
		return fmt.Errorf("error creating consul client: %w", err)
	}

	_, err = consulClient.Status().Leader()
	if err != nil {
		return fmt.Errorf("error checking consul connection status: %w", err)
	}

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
	app.Use(recover.New())

	Register(app)

	app.Listen(":3000")
	return nil
}
