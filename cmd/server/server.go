package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/marco-souza/gx/internal/config"
	"github.com/marco-souza/gx/internal/routes"
)

type Server struct {
	addr     string
	hostname string
	port     string
	app      *fiber.App
}

var conf = config.Load()

func New() *Server {
	hostname := conf.Hostname
	port := conf.Port
	addr := fmt.Sprintf("%s:%s", hostname, port)

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