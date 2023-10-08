package pages

import "github.com/marco-souza/gx/internal/config"

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
	"Marco.labs 🚀",
	"https://github.com/marco-souza/gx",
	[]MenuItem{
		{"/", "Home"},
		{"https://marco.deno.dev/blog", "Blog"},
		{"/playground", "Playground"},
	},
}
