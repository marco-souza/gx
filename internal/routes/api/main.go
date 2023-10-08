package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/entities"
)

var registers = []entities.Register{
	now,
}

func Apply(router fiber.Router) {
	for _, register := range registers {
		register(router)
	}
}
