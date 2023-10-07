package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/routes"
)

type Server struct {
	addr     string
	hostname string
	port     int
	app      *fiber.App
}

func New(hostname string, port int) *Server {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	return &Server{
		addr:     addr,
		port:     port,
		hostname: hostname,
		app:      fiber.New(),
	}
}

func (s *Server) Start() {
	s.setupRoutes()

	err := s.app.Listen(s.addr)
	if err != nil {
		fmt.Println("an error occured: ", err)
	}
}

func (s *Server) setupRoutes() {
	routes.SetupRoutes(s.app)
}
