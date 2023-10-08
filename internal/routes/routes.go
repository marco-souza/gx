package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/routes/api"
	"github.com/marco-souza/gx/internal/routes/pages"
)

func SetupRoutes(app *fiber.App) {
	pages.Apply(app.Group("/"))
	api.Apply(app.Group("/api"))
}
