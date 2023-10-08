package pages

import "github.com/gofiber/fiber/v2"

type MenuItem struct {
	Href string
	Name string
}

type Params struct {
	Title     string
	Repo      string
	MenuItems []MenuItem
}

func rootHandler(c *fiber.Ctx) error {
	params := Params{
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
