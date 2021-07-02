package app

import "github.com/gofiber/fiber/v2"

func _IndexHandler(c *fiber.Ctx) error {
	return c.Render("templates/index", &fiber.Map{
		"title": "Embassy",
		"groups": []fiber.Map{
			{
				"title": "Monitoring",
				"services": []fiber.Map{
					{
						"title":       "Grafana",
						"description": "Metrics, analytics, & dashboards",
						"url":         "http://dashboard.internal.bootleg.technology",
						"icon_url":    "https://simpleicons.org/icons/grafana.svg",
					},
					{
						"title":       "Prometheus",
						"description": "Metrics aggregation & querying",
						"url":         "http://prometheus.internal.bootleg.technology",
						"icon_url":    "https://simpleicons.org/icons/prometheus.svg",
					},
				},
			},
			{
				"title": "Management",
				"services": []fiber.Map{
					{
						"title":       "Consul",
						"description": "Service discovery & key-value store",
						"url":         "http://consul.internal.bootleg.technology",
						"icon_url":    "https://simpleicons.org/icons/consul.svg",
					},
					{
						"title":       "Nomad",
						"description": "Job scheduler",
						"url":         "http://nomad.internal.bootleg.technology",
					},
				},
			},
			{
				"title": "Storage",
				"services": []fiber.Map{
					{
						"title":       "TrueNAS",
						"description": "Network-attached storage",
						"url":         "http://192.168.1.106",
						"icon_url":    "https://simpleicons.org/icons/truenas.svg",
					},
				},
			},
			{
				"title": "Networking",
				"services": []fiber.Map{
					{
						"title":       "Traefik (Internal)",
						"description": "In-cluster HTTP request routing",
						"url":         "http://traefik.internal.bootleg.technology",
						"icon_url":    "https://symbols.getvecta.com/stencil_98/35_traefik-icon.290dcd6a8f.svg",
					},
					{
						"title":       "Traefik (External)",
						"description": "Internet-facing HTTP request routing",
						"url":         "http://traefik-external.internal.bootleg.technology",
						"icon_url":    "https://symbols.getvecta.com/stencil_98/35_traefik-icon.290dcd6a8f.svg",
					},
				},
			},
			{
				"title": "Security",
				"services": []fiber.Map{
					{
						"title":       "Vault",
						"description": "Secrets management",
						"url":         "http://vault.internal.bootleg.technology",
						"icon_url":    "https://simpleicons.org/icons/vault.svg",
					},
				},
			},
		},
	})
}

func Register(app *fiber.App) {
	app.Get("/", _IndexHandler)
}
