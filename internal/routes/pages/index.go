package pages

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/config"
)

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

func rootHandler(c *fiber.Ctx) error {
	IsDev := conf.Env == "development"

	params := Params{
		IsDev,
		"Hello, World 👋!",
		"https://github.com/marco-souza/gx",
		[]MenuItem{
			{"/", "Home"},
			{"https://marco.deno.dev/blog", "Blog"},
			{"https://marco.deno.dev/playground", "Playground"},
		},
	}

	return c.Render("index", params)
}

func root(router fiber.Router) {
	router.Get("/", rootHandler)
}
