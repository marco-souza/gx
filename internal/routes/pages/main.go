package pages

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/entities"
)

var registers = []entities.Register{
	root, playground,
	notFound,
}

func Apply(router fiber.Router) {
	for _, register := range registers {
		register(router)
	}
}
