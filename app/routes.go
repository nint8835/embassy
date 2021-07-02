package app

import "github.com/gofiber/fiber/v2"

func _IndexHandler(c *fiber.Ctx) error {
	groups := fetchServices()
	return c.Render("templates/index", &fiber.Map{
		"Title":  "Embassy",
		"Groups": groups,
	})
	// return c.Render("templates/index", &fiber.Map{
	// 	"Title": "Embassy",
	// 	"Groups": []Group{
	// 		{
	// 			Title: "Monitoring",
	// 			Services: []*Service{
	// 				{
	// 					Title:       "Grafana",
	// 					Description: "Metrics, analytics, & dashboards",
	// 					URL:         "http://dashboard.internal.bootleg.technology",
	// 					IconURL:     "https://simpleicons.org/icons/grafana.svg",
	// 					Status:      "error",
	// 				},
	// 				{
	// 					Title:       "Prometheus",
	// 					Description: "Metrics aggregation & querying",
	// 					URL:         "http://prometheus.internal.bootleg.technology",
	// 					IconURL:     "https://simpleicons.org/icons/prometheus.svg",
	// 					Status:      "warning",
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Title: "Management",
	// 			Services: []*Service{
	// 				{
	// 					Title:       "Consul",
	// 					Description: "Service discovery & key-value store",
	// 					URL:         "http://consul.internal.bootleg.technology",
	// 					IconURL:     "https://simpleicons.org/icons/consul.svg",
	// 				},
	// 				{
	// 					Title:       "Nomad",
	// 					Description: "Job scheduler",
	// 					URL:         "http://nomad.internal.bootleg.technology",
	// 					Status:      "error",
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Title: "Storage",
	// 			Services: []*Service{
	// 				{
	// 					Title:       "TrueNAS",
	// 					Description: "Network-attached storage",
	// 					URL:         "http://192.168.1.106",
	// 					IconURL:     "https://simpleicons.org/icons/truenas.svg",
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Title: "Networking",
	// 			Services: []*Service{
	// 				{
	// 					Title:       "Traefik (Internal)",
	// 					Description: "In-cluster HTTP request routing",
	// 					URL:         "http://traefik.internal.bootleg.technology",
	// 					IconURL:     "https://symbols.getvecta.com/stencil_98/35_traefik-icon.290dcd6a8f.svg",
	// 				},
	// 				{
	// 					Title:       "Traefik (External)",
	// 					Description: "Internet-facing HTTP request routing",
	// 					URL:         "http://traefik-external.internal.bootleg.technology",
	// 					IconURL:     "https://symbols.getvecta.com/stencil_98/35_traefik-icon.290dcd6a8f.svg",
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Title: "Security",
	// 			Services: []*Service{
	// 				{
	// 					Title:       "Vault",
	// 					Description: "Secrets management",
	// 					URL:         "http://vault.internal.bootleg.technology",
	// 					IconURL:     "https://simpleicons.org/icons/vault.svg",
	// 				},
	// 			},
	// 		},
	// 	},
	// })
}

func Register(app *fiber.App) {
	app.Get("/", _IndexHandler)
}
