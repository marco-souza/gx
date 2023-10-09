package pages

import (
	"log"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/github"
)

type rootProps struct {
	Params
	PrimaryBtn   string
	SecondaryBtn string
	Profile      github.GitHubUser
}

func contactURL() string {
	q := url.Values{}
	q.Add("subject", "Hi Marco, Let's have a coffee")

	// mailto does not work with spaces as '+'
	contact := "mailto:marco@tremtec.com?" + strings.ReplaceAll(
		q.Encode(), "+", "%20",
	)

	log.Println("Contact Link generated", contact)

	return contact
}

func rootHandler(c *fiber.Ctx) error {
	props := rootProps{
		PrimaryBtn:   contactURL(),
		SecondaryBtn: "/resume",
		Params:       defaultParams,
		Profile:      github.User("marco-souza"),
	}

	return c.Render("index", props)
}

func root(router fiber.Router) {
	router.Get("/", rootHandler)
}
