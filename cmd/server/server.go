package server

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/marco-souza/gx/internal/config"
	"github.com/marco-souza/gx/internal/routes"
)

type Server struct {
	addr     string
	hostname string
	port     string
	app      *fiber.App
}

var (
	conf         = config.Load()
	shouldUpdate = true
)

func New() *Server {
	hostname := conf.Hostname
	port := conf.Port
	addr := fmt.Sprintf("%s:%s", hostname, port)

	engine := html.New("./views", ".html")
	if conf.Env == "development" {
		engine.Debug(true)
		engine.Reload(true)
	}

	return &Server{
		addr:     addr,
		port:     port,
		hostname: hostname,
		app: fiber.New(fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/main",
		}),
	}
}

func (s *Server) Start() {
	s.setupRoutes()

	log.Fatal(s.app.Listen(s.addr))
}

func (s *Server) setupRoutes() {
	if conf.Env == "development" {
		log.Println("setup hot reload")
		s.app.Get("/hot-reload", hotReloadHandler)
	}

	log.Println("setup static resources")
	s.app.Static("/static", "./static", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	routes.SetupRoutes(s.app)
}

func hotReloadHandler(c *fiber.Ctx) error {
	defer func() { shouldUpdate = false }()
	return c.SendString(fmt.Sprintf("%t", shouldUpdate))
}
