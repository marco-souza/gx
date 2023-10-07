package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/", pageHandler)
	app.Get("/api/", apiHandler)
}

func pageHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func apiHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, API 👋!")
}
