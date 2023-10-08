package api

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func now(router fiber.Router) {
	router.Get("/now",
		func(c *fiber.Ctx) error {
			now := time.Now().String()
			log.Println("Now is", now)
			return c.Render("now", fiber.Map{"Time": now}, "layouts/empty")
		},
	)
}
