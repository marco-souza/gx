package pages

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/config"
	"github.com/marco-souza/gx/internal/entities"
)

var registers = []entities.Register{
	root, playground,
}

func Apply(router fiber.Router) {
	for _, register := range registers {
		register(router)
	}
}

type MenuItem struct {
	Href string
	Name string
}

type Params struct {
	IsEnv     bool
	Title     string
	Repo      string
	MenuItems []MenuItem
}

var conf = config.Load()

var defaultParams = Params{
	conf.Env == "development",
	"Hello, World 👋!",
	"https://github.com/marco-souza/gx",
	[]MenuItem{
		{"/", "Home"},
		{"https://marco.deno.dev/blog", "Blog"},
		{"/playground", "Playground"},
	},
}
