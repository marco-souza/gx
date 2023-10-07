package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func apiHandler(c *fiber.Ctx) error {
	// TODO: render gotmpl from /elements
	now := time.Now().String()
	return c.Render("now", fiber.Map{"Time": now}, "layouts/empty")
}
