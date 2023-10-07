package routes

import "github.com/gofiber/fiber/v2"

type WebParams struct {
	Title string
}

func uiHandler(c *fiber.Ctx) error {
	// TODO: handle /pages routing path
	params := WebParams{"Hello, World 👋!"}
	return c.Render("index", params)
}
