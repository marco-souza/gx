package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/", uiHandler)
	app.Get("/api/now", apiHandler)
}
