package pages

import "github.com/gofiber/fiber/v2"

func rootHandler(c *fiber.Ctx) error {
	return c.Render("index", defaultParams)
}

func root(router fiber.Router) {
	router.Get("/", rootHandler)
}
