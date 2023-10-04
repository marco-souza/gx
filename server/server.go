package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	host string
	port int
	app  *fiber.App
}

func New() Server {
	return Server{
		host: "",
		port: 8001,
		app:  fiber.New(),
	}
}

func (s *Server) Start() {
	s.setupRoutes()

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	err := s.app.Listen(addr)
	if err != nil {
		fmt.Println("an error occured: ", err)
	}
}

func (s *Server) setupRoutes() {
	// TODO: build page routes using <root>/routes folder
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
