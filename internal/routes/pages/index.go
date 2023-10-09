package pages

import (
	"log"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type IndexProps struct {
	Params
	PrimaryBtn   string
	SecondaryBtn string
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
	props := IndexProps{
		PrimaryBtn:   contactURL(),
		SecondaryBtn: "/resume",
		Params:       defaultParams,
	}

	return c.Render("index", props)
}

func root(router fiber.Router) {
	router.Get("/", rootHandler)
}
