package pages

import "github.com/gofiber/fiber/v2"

func playgroundHandler(c *fiber.Ctx) error {
	return c.Render("playground", defaultParams)
}

func playground(router fiber.Router) {
	router.Get("/playground", playgroundHandler)
}
